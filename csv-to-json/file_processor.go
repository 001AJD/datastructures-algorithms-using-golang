package main

const BATCH_SIZE = 100
const WORKER_POOL = 20

func ProcessFile(path string) bool {
	if _, err := ParallelProcessecor(path); err != nil {
		return false
	} else {
		return true
	}
}
