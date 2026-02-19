package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sync"
)

// outputWriter drains transformed rows from processedRows and writes them as a
// JSON array into the provided output file.
func outputWriter(processedRows <-chan Domains, errorChannel chan<- string, writer *bufio.Writer, writerWg *sync.WaitGroup) {
	defer writerWg.Done()
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
