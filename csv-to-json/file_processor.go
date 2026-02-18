package main

const BATCH_SIZE = 100
const WORKER_POOL = 20

// ProcessFile runs the CSV-to-JSON pipeline for the file at the given path.
// It returns true when processing succeeds and false when an error occurs.
func ProcessFile(path string) bool {
	if _, err := ParallelProcessecor(path); err != nil {
		return false
	} else {
		return true
	}
}
