package internal

import (
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	// Save original os.Args
	originalArgs := os.Args
	// Restore os.Args after test
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"_", "--minsize=100", "path"}

	argsProcessor := NewArgsParser()

	if argsProcessor.StartPath != "path" {
		t.Fatalf("expected 'path', got %s", argsProcessor.StartPath)
	}

	if argsProcessor.MinSize != 100 {
		t.Fatalf("expected 100, got %d", argsProcessor.MinSize)
	}
}
