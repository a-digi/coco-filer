package integration

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/a-digi/coco-filer/filer"
)

type inMemoryMultipartFile struct {
	*strings.Reader
}

func (f *inMemoryMultipartFile) Close() error { return nil }

func TestIntegration_FileManager_MoveFile(t *testing.T) {
	dir := t.TempDir()
	fm := filer.NewFileManager(dir)

	srcFile := filepath.Join(dir, "integration.txt")
	dstFile := filepath.Join(dir, "integration_subdir", "integration.txt")
	content := []byte("integration test")
	if err := os.WriteFile(srcFile, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	if err := fm.MoveFile(srcFile, dstFile); err != nil {
		t.Fatalf("MoveFile failed: %v", err)
	}

	if _, err := os.Stat(dstFile); err != nil {
		t.Errorf("Destination file not found: %v", err)
	}
	if _, err := os.Stat(srcFile); !os.IsNotExist(err) {
		t.Errorf("Source file still exists after move")
	}
}

func TestIntegration_FileMultiPartManager_MoveFile(t *testing.T) {
	dir := t.TempDir()
	fm := filer.NewFileMultiPartManager(dir)
	data := "integration multipart content"
	file := &inMemoryMultipartFile{Reader: strings.NewReader(data)}
	defer file.Close()

	header := &multipart.FileHeader{
		Filename: "integration_upload.txt",
		Header:   make(map[string][]string),
	}
	header.Header.Set("Content-Type", "text/plain")

	f, err := fm.MoveFile(file, header)
	if err != nil {
		t.Fatalf("MoveFile (multipart) failed: %v", err)
	}
	if f == nil || f.Name != "integration_upload.txt" {
		t.Errorf("Returned file struct invalid: %+v", f)
	}
	if f.Size != int64(len(data)) {
		t.Errorf("File size mismatch: got %d, want %d", f.Size, len(data))
	}
	if f.MimeType != "text/plain" {
		t.Errorf("MimeType mismatch: got %s", f.MimeType)
	}
}
