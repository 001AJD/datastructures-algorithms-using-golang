package processor

import (
	"fmt"
	"time"
)

const BATCH_SIZE = 100
const WORKER_POOL = 20

// ProcessFile runs the CSV-to-JSON pipeline for the file at the given path.
// It returns true when processing succeeds and false when an error occurs.
func ProcessFile(inputFilePath string, outputFilePath string) (bool, error) {
	start1 := time.Now()
	if result, err := ParallelProcessecor(inputFilePath, outputFilePath); err != nil {
		elapsed := time.Since(start1)
		fmt.Println("Time take by parallel processor", elapsed)
		return result, err
	} else {
		elapsed := time.Since(start1)
		fmt.Println("Time take by parallel processor", elapsed)
		return result, nil
	}

}
