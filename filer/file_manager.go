package filer

import (
	"os"
	"time"
)

type File struct {
	Name     string
	Path     string
	Size     int64
	ModTime  time.Time
	Mode     os.FileMode
	MimeType string
}

type FileManager struct {
	DefaultDir string
}

func NewFileManager(defaultDir string) *FileManager {

	if defaultDir == "" {
		defaultDir = "data/uploads"
	}

	return &FileManager{DefaultDir: defaultDir}
}
