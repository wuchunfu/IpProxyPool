package fileutil

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

// 判断所给路径是否为文件夹
// IsDir returns true if given path is a dir,
// or returns false when it's a directory or does not exist.
func IsDir(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// 判断所给路径是否为文件
// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	return !IsDir(filePath)
}

// FileExist 判断所给路径文件/文件夹是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// PathExists return true if given path exist.
func PathExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// Sha1f return file sha1 encode
func Sha1f(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// ReadFile 读取文件
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}
