package main

import (
	"testing"
)

func TestDeepCompare(t *testing.T) {
	fileProcessor := newFileProcessor()
	file := "../test/file2"
	fileToCompareDup := "../test/duplicateoffile2"
	fileToCompareNotDupSameLength := "../test/file4"

	isSame, err := fileProcessor.deepCompare(file, fileToCompareDup)
	if !isSame || err != nil {
		t.Fatalf("expected true, got %t, %v", isSame, err)
	}

	isSame, err = fileProcessor.deepCompare(file, fileToCompareNotDupSameLength)
	if isSame || err != nil {
		t.Fatalf("expected false, got %t, %v", isSame, err)
	}
}
