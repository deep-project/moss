package osx

import (
	"embed"
	"errors"
	"path/filepath"
)

func EmbedCopyToDir(fs embed.FS, memDir, localDir string) error {
	files, err := fs.ReadDir(filepath.ToSlash(memDir))
	if err != nil {
		return errors.New("read dir failed." + err.Error())
	}
	for _, f := range files {
		if f.IsDir() {
			if err := EmbedCopyToDir(fs, filepath.Join(memDir, f.Name()), filepath.Join(localDir, f.Name())); err != nil {
				return err
			}
			continue
		}
		if err := EmbedCopyFileToLocal(fs, filepath.Join(memDir, f.Name()), filepath.Join(localDir, f.Name())); err != nil {
			return err
		}
	}
	return nil
}

func EmbedCopyFileToLocal(fs embed.FS, memFile, localFile string) error {
	b, err := fs.ReadFile(filepath.ToSlash(memFile))
	if err != nil {
		return errors.New("read file failed. " + err.Error())
	}
	if err = CreateFileWithDir(localFile, b, 0777); err != nil {
		return errors.New("create file failed." + err.Error())
	}
	return nil
}
