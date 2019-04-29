package util

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 获取当前可执行文件名（不含扩展名）
func GetProgramFileBaseName() string {
	_, file := filepath.Split(os.Args[0])
	if len(file) > 0 {
		ext := path.Ext(file)
		if len(ext) > 0 {
			file = strings.TrimSuffix(file, ext)
		}
	}
	return file
}

// 如果目录不存在创建指定目录
func IsDirOrFileExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		panic(err)
	}
	return nil
}

// 如果目录不存在创建指定目录
func MustMkdirIfNotExist(path string) {
	if err := IsDirOrFileExist(path); err != nil {
		os.MkdirAll(path, os.ModePerm)
	}
}
