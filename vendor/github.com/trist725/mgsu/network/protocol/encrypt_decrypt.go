package protocol

type IEncryptor interface {
	Encrypt(data []byte) (err error)
}

type EncryptFunc func(data []byte) (err error)

func (fn EncryptFunc) Encrypt(data []byte) error {
	return fn(data)
}

func DefaultEncryptor(data []byte) error {
	size := len(data)
	for i := 0; i < size; i++ {
		data[i] = byte(int(data[i]) ^ (size - i))
	}
	return nil
}

type IDecryptor interface {
	Decrypt(data []byte) (err error)
}

type DecryptFunc func(data []byte) (err error)

func (fn DecryptFunc) Decrypt(data []byte) error {
	return fn(data)
}

func DefaultDecryptor(data []byte) error {
	size := len(data)
	for i := 0; i < size; i++ {
		data[i] = byte(int(data[i]) ^ (size - i))
	}
	return nil
}

type NonEncryptor struct {
}

func (NonEncryptor) Encrypt([]byte) error {
	return nil
}

type NonDecryptor struct {
}

func (NonDecryptor) Decrypt([]byte) error {
	return nil
}
