package osx

import (
	"errors"
	"os"
	"path/filepath"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func CreateDirIsNotExist(dir string, perm os.FileMode) (err error) {
	if IsExist(dir) {
		return
	}
	return os.MkdirAll(dir, perm)
}

func CreateFileWithDir(path string, content []byte, perm os.FileMode) (err error) {
	if err = CreateDirIsNotExist(filepath.Dir(path), perm); err != nil {
		return
	}
	return CreateFile(path, content)
}

func CreateFile(path string, content []byte) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	if len(content) > 0 {
		_, err = file.Write(content)
	}
	return
}
