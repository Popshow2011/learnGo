package storage

import (
	"os"
)

type Storage struct {
	FilePath string
}

func (s Storage) SaveBinToJson(bin []byte) (bool, error) {
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

func (s Storage) ReadBinToJson() ([]byte, error) {
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

func NewStorage(path string) *Storage {
	return &Storage{
		FilePath: path,
	}
}
