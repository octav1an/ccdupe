package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	listDir("./test")
}

var hashes = make(map[[16]byte]string)

func listDir(folder string) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			// Recursively read the files in subdirectories
			listDir(filepath.Join(folder, e.Name()))
		} else {
			var path = filepath.Join(folder, e.Name())

			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			compareByHash(path, data)
		}
	}

	fmt.Println(hashes)
}

func compareByHash(path string, data []byte) {
	var hash = md5.Sum(data)

	existing_path, exists := hashes[hash]
	if !exists {
		hashes[hash] = path
		fmt.Printf("No file with this hash %s and path %s exists\n", hex.EncodeToString(hash[:]), path)
	} else {
		fmt.Printf("Duplicate for %s is in %s\n", path, existing_path)
	}
}
