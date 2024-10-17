package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	listDir("./test")
}

var sizes = make(map[int64]string)

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
			// fmt.Println(filepath.Join(folder, e.Name()))
			var path = filepath.Join(folder, e.Name())
			fi, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			compareBySize(path, fi)
		}
	}
}

func compareBySize(file_path string, fi os.FileInfo) {
	original_path, ok := sizes[fi.Size()]
	if ok {
		fmt.Printf("Potential duplicates: %s %s \n", file_path, original_path)
	} else {
		sizes[fi.Size()] = file_path
	}
}
