package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error // interface{} accepts any kind of data
	Write(data interface{}) error
}

type Type string // Force the use of Type instead of any string when calling the file

const (
	FileType Type = "file"
)

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func NewStore(store Type, filename string) Store {
	switch store {
	case FileType:
		return &fileStore{filename}
	}
	return nil
}
