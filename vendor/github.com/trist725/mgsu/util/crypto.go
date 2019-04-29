package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"sync"
)

var (
	md5Hasher     = md5.New()
	md5HasherLock = &sync.Mutex{}
)

// 生成md串
func MD5Sum(input []byte) (output []byte) {
	md5HasherLock.Lock()
	defer md5HasherLock.Unlock()

	output = md5Hasher.Sum(input)

	return
}

// 生成md串
func MD5Sumf(format string, args ...interface{}) (output []byte) {
	md5HasherLock.Lock()
	defer md5HasherLock.Unlock()

	input := []byte(fmt.Sprintf(format, args...))

	output = md5Hasher.Sum(input)

	return
}

// rsa公钥加密
func RsaPubEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// rsa私钥解密
func RsaPrivDecrypt(privateKey []byte, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, origData)
}
