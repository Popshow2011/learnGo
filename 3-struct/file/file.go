package file

import (
	"encoding/json"
	"os"
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

func (f File) IsJson(data []byte) bool {
	return json.Valid(data)
}

func NewConstructor(path string) *File {
	return &File{
		Path: path,
	}
}
