package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
			fp.compareByHash(path)
		}
	}
	return nil
}

func (fp *FileProcessor) compareByHash(path string) {
	h, _ := fp.calculateHash(path)
	// Convert the hash to a 16 byte array
	var hash [16]byte
	copy(hash[:], h[:16])

	existingPath, exists := fp.hashes[hash]
	if exists {
		// If hash dup exists compare every byte
		isByteDup, err := fp.deepCompare(path, existingPath)
		if err != nil {
			fmt.Errorf("error deep compare for file %s: %w", path, err)
		}

		if isByteDup {
			fmt.Errorf("duplicate for %s is in %s\n", path, existingPath)
		}
	} else {
		fp.hashes[hash] = path
	}
}

// Calculate md5 hash for a given file path. File is read in chunks
func (fp *FileProcessor) calculateHash(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	hash := md5.New()
	buf := make([]byte, 4096)
	for {
		bytesCount, err := file.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
		}

		if bytesCount == 0 {
			break
		}

		if _, err := hash.Write(buf[:bytesCount]); err != nil {
			return nil, fmt.Errorf("error writing hash: %w", err)
		}
	}
	return hash.Sum(nil), nil
}

// Compare the bytes in data against file from compareToPath
// TODO: refactor to read in chunks
func (fp *FileProcessor) deepCompare(filePath string, compareToFilePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	fileToCompare, err := os.Open(compareToFilePath)
	if err != nil {
		return false, fmt.Errorf("error opening file %s: %w", compareToFilePath, err)
	}
	defer fileToCompare.Close()

	// Compare bytes
	bufFile := make([]byte, 4096)
	bufFileToCompare := make([]byte, 4096)
	isDup := false

	for {
		fileBytesCount, err := file.Read(bufFile)
		if err != nil && err.Error() != "EOF" {
			return false, fmt.Errorf("error reading file %s: %w", filePath, err)
		}

		fileToCompareBytesCount, err := fileToCompare.Read(bufFileToCompare)
		if err != nil && err.Error() != "EOF" {
			return false, fmt.Errorf("error reading file %s: %w", compareToFilePath, err)
		}

		if fileBytesCount != fileToCompareBytesCount {
			break
		} else if fileBytesCount == 0 && fileToCompareBytesCount == 0 {
			isDup = true
			break
		}

		if !bytes.Equal(bufFile, bufFileToCompare) {
			return false, nil
		}
	}

	return isDup, nil
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	printMemUsage()
	fileProcessor := newFileProcessor()
	if err := fileProcessor.ProcessDirectory("../test"); err != nil {
		fmt.Println("error:", err)
	}
	printMemUsage()
}
