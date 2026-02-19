package main

import (
	"fmt"
	"time"
)

const BATCH_SIZE = 100
const WORKER_POOL = 20

// ProcessFile runs the CSV-to-JSON pipeline for the file at the given path.
// It returns true when processing succeeds and false when an error occurs.
func ProcessFile(path string) bool {
	start1 := time.Now()
	var parallelProcessorResult bool
	if result, err := ParallelProcessecor(path); err != nil {
		parallelProcessorResult = result
	} else {
		parallelProcessorResult = result
	}
	elapsed := time.Since(start1)
	fmt.Println("Time take by parallel processor", elapsed)
	return parallelProcessorResult
}
