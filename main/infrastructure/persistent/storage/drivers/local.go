package drivers

import (
	"errors"
	"io"
	"moss/infrastructure/persistent/storage"
	"moss/infrastructure/utils/osx"
	"os"
	"path/filepath"
)

type Local struct {
	Path string `json:"path"`
}

func (l *Local) Init() error {
	if l.Path == "" {
		return errors.New("path undefined")
	}
	return osx.CreateDirIsNotExist(l.Path, 0777)
}

func (l *Local) Close() error {
	return nil
}

func (l *Local) Get(key string) (*storage.GetValue, error) {
	if l.Path == "" {
		return nil, errors.New("path undefined")
	}
	f, err := os.Open(l.filePath(key))
	if err != nil {
		return nil, err
	}
	return storage.NewGetValue(f), nil
}

func (l *Local) Set(key string, val *storage.SetValue) (err error) {
	if l.Path == "" {
		return errors.New("path undefined")
	}
	path := l.filePath(key)
	_ = l.createDir(path)

	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = io.Copy(f, val.Reader)
	return
}

func (l *Local) Delete(key string) error {
	if l.Path == "" {
		return errors.New("path undefined")
	}
	path := l.filePath(key)
	if path == "" || path == "/" {
		return errors.New("path is empty or root dir")
	}
	return os.Remove(path)
}

func (l *Local) createDir(path string) error {
	dir, _ := filepath.Split(path)
	return osx.CreateDirIsNotExist(dir, 0777)
}

func (l *Local) filePath(key string) string {
	return filepath.Join(l.Path, key)
}

//func (l *Local) originalPath(key string) string {
//	path, name := filepath.Split(key)
//	// 给文件名加具体扩展，可以防止文件和目录冲突
//	// 比如 /a/b/1 和 /a/b/1/2 ; 或者 /a/b.html 和 /a/b.html/2
//	// 如果已经创建了第一个文件，那创建第二个目录会失败
//	return filepath.Join(path, name+".file")
//}
//
//func (l *Local) hashPath(key string) string {
//	hashKey := cryptor.Md5String(key)
//	return filepath.Join(hashKey[0:2], hashKey[2:4], hashKey[4:6], hashKey[6:])
//}
