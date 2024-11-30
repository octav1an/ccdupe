package main

import (
	"fmt"

	"ccdupe/internal"
)

func main() {
	internal.PrintMemUsage()
	argsParser := internal.NewArgsParser()
	fmt.Println(argsParser.StartPath, argsParser.MinSize)

	fileProcessor := internal.NewFileProcessor()
	if err := fileProcessor.ProcessDirectory(argsParser.StartPath); err != nil {
		fmt.Println("error:", err)
	}
	internal.PrintMemUsage()
}
