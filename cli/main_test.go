package main

import (
	"path/filepath"
	"testing"
)

var fixturePath = "../fixtures/"

func TestDeepCompare(t *testing.T) {
	fileProcessor := newFileProcessor()
	file := filepath.Join(fixturePath, "file2")
	fileToCompareDup := filepath.Join(fixturePath, "duplicateoffile2")
	fileToCompareNotDupSameLength := filepath.Join(fixturePath, "file4")

	isSame, err := fileProcessor.deepCompare(file, fileToCompareDup)
	if !isSame || err != nil {
		t.Fatalf("expected true, got %t, %v", isSame, err)
	}

	isSame, err = fileProcessor.deepCompare(file, fileToCompareNotDupSameLength)
	if isSame || err != nil {
		t.Fatalf("expected false, got %t, %v", isSame, err)
	}
}
