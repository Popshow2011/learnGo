package storage

import (
	"os"
)

type FileStorage struct {
	FilePath string
}

type Storage interface {
	Write([]byte) (bool, error)
	Read() ([]byte, error)
}

func (s FileStorage) Write(bin []byte) (bool, error) {
	_, err := os.OpenFile(s.FilePath, 0, 0)
	if err != nil {
		return false, err
	}

	err = os.WriteFile(s.FilePath, bin, 0)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s FileStorage) Read() ([]byte, error) {
	_, err := os.OpenFile(s.FilePath, 0, 0)
	if err != nil {
		return nil, err
	}
	bytes, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NewStorage(path string) *FileStorage {
	return &FileStorage{
		FilePath: path,
	}
}
