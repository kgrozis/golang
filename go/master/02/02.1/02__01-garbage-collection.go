package main

import (
	"fmt"
	"runtime"
	"time"
)

// call runtime.ReadMemstats() to get recent garbage collection stats
func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 500000) // create multiple slices to allocate memory and trigger garbage collection
		if s == nil {
			fmt.Println("Operation Failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 10000000)
		if s == nil {
			fmt.Println("Operation Failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)

}
