package processor

import (
	"csv-to-json/internal/filesystem"
	"csv-to-json/internal/logging"
	"encoding/csv"
	"fmt"
	"io"
	"sync"
)

// ParallelProcessecor reads a CSV file, transforms each row into a Domains
// value in parallel, and writes the resulting JSON array to output/output.json.
func ParallelProcessecor(inputFilePath string, outputFilePath string) (bool, error) {
	const WORKER_POOL = 10
	const BATCH_SIZE = 200
	var workerWg sync.WaitGroup
	var writerWg sync.WaitGroup
	var loggerWg sync.WaitGroup
	jobs := make(chan []string, BATCH_SIZE)  // channel to read, transform the CSV row
	processedRows := make(chan Domains, 200) // channel to write the results into the output json file
	errorChannel := make(chan string, 200)   // channel to log errors

	inputFile, err := filesystem.NewBufferedFileReader(inputFilePath)
	if err != nil {
		return false, fmt.Errorf("error occurred when reading file %w", err)
	}
	defer inputFile.Close()

	outputFile, err := filesystem.NewBufferedFileWriter(outputFilePath)
	if err != nil {
		return false, err
	}
	defer outputFile.Close()

	// initialize logger, fail if logger initialization fails
	if err := logging.InitializeLogger(); err != nil {
		return false, err
	}

	// loop to spin up worker pool, to transfor the data into json
	for range WORKER_POOL {
		workerWg.Add(1)
		go worker(jobs, processedRows, errorChannel, &workerWg)
	}

	func() {
		loggerWg.Add(1)
		go logging.ProcessLogs(errorChannel, &loggerWg)

	}()

	go func() {
		workerWg.Wait()
		close(processedRows)
	}()

	// spin up a go routine to write the processed json objects into the output json file
	writerWg.Add(1)
	go outputWriter(processedRows, errorChannel, outputFile.BufferedWriter, &writerWg)

	// start processing csv file
	reader := csv.NewReader(inputFile.BufferedReader)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println("End of file reached")
			break
		}
		if err != nil {
			fmt.Println("Error occured while reading row from csv file", err.Error())
		}
		jobs <- record
	}
	close(jobs)
	writerWg.Wait()
	close(errorChannel)
	loggerWg.Wait()
	return true, nil
}
