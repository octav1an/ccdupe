package internal

import (
	"flag"
)

type ArgsParser struct {
	StartPath string
	MinSize   int // In KB
}

func NewArgsParser() *ArgsParser {
	ap := &ArgsParser{}
	ap.ParseArgs()
	return ap
}

func (ap *ArgsParser) ParseArgs() {
	minSize := flag.Int("minsize", 0, "Minimum file size to process")
	flag.Parse()

	ap.MinSize = *minSize

	// The last arg is always the path to process
	tailArgs := flag.Args()
	ap.StartPath = tailArgs[len(tailArgs)-1]
}
