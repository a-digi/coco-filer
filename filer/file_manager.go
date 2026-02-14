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

// FileManager kapselt Dateioperationen wie Speichern und Verschieben.
type FileManager struct {
	DefaultDir string
}

// NewFileManager erstellt einen neuen FileManager mit Standardverzeichnis (z.B. ./uploads).
func NewFileManager(defaultDir string) *FileManager {

	if defaultDir == "" {
		defaultDir = "data/uploads"
	}

	return &FileManager{DefaultDir: defaultDir}
}
