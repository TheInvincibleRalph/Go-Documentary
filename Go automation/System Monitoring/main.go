package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys - %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func monitorSystem() {
	for {
		printMemUsage()
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go monitorSystem()
}
