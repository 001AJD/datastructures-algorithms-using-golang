package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

var BATCH_SIZE = 500   // This is now the size of the jobs channel buffer
var WORKER_COUNT = 100 // The number of concurrent worker goroutines

// ProcessFile reads a CSV file line by line and processes its first column (domains) concurrently
// using a worker pool. It utilizes a sync.WaitGroup to ensure all workers finish before the function returns.
//
// Workflow:
//  1. Opens the file at the specified filePath.
//  2. Creates a buffered channel `jobs` to send domains to workers.
//  3. Spawns a fixed number of `WORKER_COUNT` goroutines (workers).
//  4. Reads the CSV file record by record and sends each domain into the `jobs` channel.
//  5. After reading the entire file, it closes the `jobs` channel.
//  6. Blocks at wg.Wait() until all worker goroutines have finished processing all jobs and returned.
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

// worker continuously consumes domain strings from the provided jobs channel
// and performs status verification using CheckStatus(url).
//
// Behavior:
//   - The function runs as a long-lived goroutine.
//   - It blocks on the jobs channel and processes incoming URLs sequentially.
//   - It exits gracefully when the jobs channel is closed.
//   - Upon exit, it signals completion to the provided WaitGroup.
//
// Concurrency Model:
//   - Multiple worker instances can be launched to achieve bounded parallelism.
//   - The total number of active workers determines the global concurrency level.
//   - The jobs channel acts as a coordination mechanism between producer
//     (domain feeder) and consumers (workers).
//
// Parameters:
//   - jobs: A read-only channel that supplies domain strings for processing.
//   - wg: A pointer to sync.WaitGroup used to track worker lifecycle.
//
// Note:
//   - This worker processes jobs sequentially per goroutine.
//   - Overall throughput scales with the number of worker goroutines spawned.
//   - Proper channel closure is required to allow workers to terminate cleanly.
func worker(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		CheckStatus(url)
	}
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
//  3. Launches an anonymous goroutine for each domain to perform a HEAD request via CheckStatus.
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
		go func(url string) { // go routines are spawned for a HEAD request for each domain, currently number of goroutines = number of domains, bad design, need to refactor
			defer wg.Done()
			CheckStatus(url)
		}(v)
	}
	wg.Wait()
}
