package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("URL status checker initiating...")
	PrintMemUsage()
	start := time.Now()
	ProcessFile("./data/top-100.csv")
	elapsed := time.Since(start)
	fmt.Printf("Processing completed in %s \n", elapsed)
	PrintMemUsage()
	time.Sleep(10 * time.Second)
	fmt.Println("Goroutines 10 second after process completes ", runtime.NumGoroutine())
}
