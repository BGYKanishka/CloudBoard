package storage

import (
	"io"
	"os"
	"path/filepath"
)

type Storage interface {
	Upload(filename string, file io.Reader) (string, error)
}

type LocalStorage struct {
	BaseDir string
}

func NewLocalStorage(baseDir string) *LocalStorage {
	os.MkdirAll(baseDir, 0755)
	return &LocalStorage{BaseDir: baseDir}
}

func (s *LocalStorage) Upload(filename string, file io.Reader) (string, error) {
	destPath := filepath.Join(s.BaseDir, filename)
	out, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}

	return "/uploads/" + filename, nil
}
