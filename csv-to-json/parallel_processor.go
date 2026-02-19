package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

// ParallelProcessecor reads a CSV file, transforms each row into a Domains
// value in parallel, and writes the resulting JSON array to output/output.json.
func ParallelProcessecor(filePath string) (bool, error) {
	var result bool = false
	const WORKER_POOL = 1
	const BATCH_SIZE = 200
	var wg sync.WaitGroup

	// open file to read csv file
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("Error %s occured while reading file from path %s", err.Error(), filePath)
	}
	defer file.Close()

	// create file to write data to the output file
	outputFile, err := os.OpenFile("./output/output.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to create the output file")
		return false, fmt.Errorf("Error occured while creating a output file :: %s", err.Error())
	}
	defer outputFile.Close()

	// create logger file
	logFile, err := os.OpenFile("./logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to create the log file, halting!")
		return false, fmt.Errorf("Error occured while create a log file %w", err)
	}
	defer logFile.Close()

	reader := csv.NewReader(file)

	// channel to read, transform the CSV row
	jobs := make(chan []string, BATCH_SIZE)

	// channel to write the results into the output json file
	processedRows := make(chan Domains, 200)

	// channel to log errors
	errorChannel := make(chan string, 200)

	// loop to spin up worker pool, to transfor the data into json
	for range WORKER_POOL {
		wg.Add(1)
		go worker(jobs, processedRows, errorChannel, &wg)
	}

	// spin up go routine to write logs into log file via error Channel
	// var loggerWg sync.WaitGroup
	// loggerWg.Add(1)
	go processLogs(errorChannel, logFile)

	go func() {
		wg.Wait()
		close(processedRows)
		close(errorChannel)
	}()

	// spin up a go routine to write the processed json objects into the output json file
	var writerWg sync.WaitGroup
	writerWg.Add(1)
	go outputWriter(processedRows, errorChannel, outputFile, &writerWg)

	// loop to iterate over CSV file
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
	result = true
	return result, nil
}

// worker consumes CSV records from jobs, converts each record to a Domains
// object, and sends the transformed value to processedRows.
func worker(jobs <-chan []string, processedRows chan<- Domains, errorChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range jobs {
		if outputJson, err := CSVToJSON(record); err != nil {
			// put error message in the error channel
			errorChannel <- err.Error()
		} else {
			processedRows <- outputJson
		}
	}
}

// outputWriter drains transformed rows from processedRows and writes them as a
// JSON array into the provided output file.
func outputWriter(processedRows <-chan Domains, errorChannel chan<- string, outputFile *os.File, writerWg *sync.WaitGroup) {
	defer writerWg.Done()
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	if _, err := writer.WriteString("[\n"); err != nil {
		errorChannel <- err.Error()
		fmt.Println("failed to write to the output file")
	}

	first := true
	for p := range processedRows {
		if !first {
			if _, err := writer.WriteString(",\n"); err != nil {
				errorChannel <- err.Error()
				fmt.Println("failed to write")
			}
		}

		bytes, err := json.Marshal(p)
		if err != nil {
			errorChannel <- err.Error()
			fmt.Println("Error occured during json marshal")
		}
		if _, err := writer.Write(bytes); err != nil {
			errorChannel <- err.Error()
			fmt.Println("failed to write")
		}
		first = false
	}
	if _, err := writer.WriteString("\n]"); err != nil {
		errorChannel <- err.Error()
		fmt.Println("failed to write")
	}
}
