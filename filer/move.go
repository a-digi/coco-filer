package filer

import (
	"fmt"
	"os"
	"path/filepath"
)

func (fm *FileManager) MoveFile(src, dst string) error {
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("konnte Zielverzeichnis nicht anlegen: %w", err)
	}
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("Fehler beim Verschieben: %w", err)
	}
	return nil
}