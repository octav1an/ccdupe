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

			fmt.Println(e.Name())
		}
	}
}
