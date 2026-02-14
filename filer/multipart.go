package filer

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Factory-Funktion für FileMultiPartManager
func NewFileMultiPartManager(defaultDir string) *FileMultiPartManager {
	return &FileMultiPartManager{FileManager: NewFileManager(defaultDir)}
}

type FileMultiPartManager struct {
    FileManager *FileManager
}

func (fm *FileMultiPartManager) MoveFile(file multipart.File, handler *multipart.FileHeader, targetDir ...string) (*File, error) {
	dir := fm.FileManager.DefaultDir

	if len(targetDir) > 0 && targetDir[0] != "" {
		dir = targetDir[0]
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("konnte Zielverzeichnis nicht anlegen: %w", err)
	}

	filePath := filepath.Join(dir, handler.Filename)
	out, err := os.Create(filePath)

	if err != nil {
		return nil, fmt.Errorf("konnte Datei nicht erstellen: %w", err)
	}

	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return nil, fmt.Errorf("Fehler beim Schreiben der Datei: %w", err)
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen der Dateiinformationen: %w", err)
	}

	f := &File{
		Name:     info.Name(),
		Path:     filePath,
		Size:     info.Size(),
		ModTime:  info.ModTime(),
		Mode:     info.Mode(),
		MimeType: handler.Header.Get("Content-Type"),
	}

	return f, nil
}
