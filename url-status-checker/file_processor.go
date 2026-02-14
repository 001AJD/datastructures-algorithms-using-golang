// TODO: currently the goroutines are unbounded and directly depends on the number of urls in the file
// example if there are 10,000 urls --> 10,000 goroutines are spawned, that is bad and needs be refactored
// what's next ? - The concurrency should be bounded irrespective of the input size

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

var BATCH_SIZE = 500   // bigger batch size can consume more memory, increasing this does not make processing faster
var WORKER_COUNT = 100 // check the file descriptor size and decide the worker count

// ProcessFile reads a CSV file line by line and processes its first column (domains) in concurrent batches.
// It utilizes a sync.WaitGroup to ensure all background goroutines finish processing before the function returns.
//
// Workflow:
//  1. Opens the file at the specified filePath.
//  2. Iterates through the CSV records, accumulating domain strings into a slice (batch).
//  3. Once the batch reaches BATCH_SIZE, it launches a goroutine to execute ProcessBatch(batch).
//  4. Handles the final partial batch if the file ends before reaching a full BATCH_SIZE.
//  5. Blocks at wg.Wait() until all asynchronous processing is complete.
//
// Parameters:
//   - filePath: The local system path to the CSV file to be processed.
//
// Note: This function will panic if the file cannot be opened or if a CSV read error
// other than io.EOF occurs.
func ProcessFile(filePath string) {
	var wg sync.WaitGroup
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jobs := make(chan string, BATCH_SIZE) // buffered channel

	for i := 0; i < WORKER_COUNT; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		domain := record[0] // first column in csv is domain
		jobs <- domain

	}
	close(jobs)
	wg.Wait()
	fmt.Printf("\nFile processing complete\n")
}

// ProcessBatch handles a collection of domain strings by initiating concurrent status checks.
// It iterates through the provided batch and spawns a separate goroutine for each domain
// to execute CheckStatus(url) in parallel.
//
// Parameters:
//   - batch: A slice of strings containing the domains or URLs to be processed.
//
// Workflow:
//  1. Prints the total count of items in the current batch for logging.
//  2. Iterates over the slice, adding to a sync.WaitGroup for each entry.
//  3. Launches an anonymous goroutine for each domain to perform a GET request via CheckStatus.
//  4. Blocks execution at wg.Wait() until every goroutine in the batch has signaled completion.
//
// Warning: This implementation spawns one goroutine per domain without a limit.
// For large batches, this can lead to high memory consumption or "too many open files"
// errors from the operating system.
func ProcessBatch(batch []string) {
	var wg sync.WaitGroup
	// fmt.Println(len(batch))
	for _, v := range batch {
		wg.Add(1)
		go func(url string) { // go routines are spawned for a GET request for each domain, currently number of goroutines = number of domains, bad design, need to refactor
			defer wg.Done()
			CheckStatus(url)
		}(v)
	}
	wg.Wait()
}

func worker(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		CheckStatus(url)
	}

}
