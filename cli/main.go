package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
)

type FileProcessor struct {
	hashes map[[16]byte]string
}

func newFileProcessor() *FileProcessor {
	return &FileProcessor{hashes: make(map[[16]byte]string)}
}

func (fp *FileProcessor) ProcessDirectory(folder string) error {
	return fp.listDir(folder)
}

func (fp *FileProcessor) listDir(folder string) error {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %w", folder, err)
	}

	for _, e := range entries {
		path := filepath.Join(folder, e.Name())
		if e.IsDir() {
			// Recursively read the files in subdirectories
			if err := fp.listDir(path); err != nil {
				return err
			}
		} else {
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}
			fp.compareByHash(path, data)
		}
	}
	return nil
}

func (fp *FileProcessor) compareByHash(path string, data []byte) {
	hash := md5.Sum(data)

	existing_path, exists := fp.hashes[hash]
	if exists {
		fmt.Printf("Duplicate for %s is in %s\n", path, existing_path)
	} else {
		fp.hashes[hash] = path
	}
}

func main() {
	fileProcessor := newFileProcessor()
	if err := fileProcessor.ProcessDirectory("./test"); err != nil {
		fmt.Println("Error:", err)
	}
}
