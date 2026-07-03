package file

import (
	"os"
	"path/filepath"
)

type File struct {
	Path string
}

func (f File) ReadFile() ([]byte, error) {
	_, err := os.OpenFile(f.Path, 0, 0)
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (f File) IsJson() bool {
	if filepath.Ext(f.Path) == ".json" {
		return true
	} else {
		return false
	}
}

func NewConstructor(path string) *File {
	return &File{
		Path: path,
	}
}
