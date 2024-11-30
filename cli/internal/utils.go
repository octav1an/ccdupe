package internal

import (
	"fmt"
	"runtime"
)

func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func BToKb(b uint64) uint64 {
	return b / 1024
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", BToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", BToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", BToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
