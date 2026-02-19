package main

import (
	"bufio"
	"fmt"
	"os"
)

func processLogs(errorChannel <-chan string, file *os.File) {
	writer := bufio.NewWriter(file)
	defer func() {
		if err := writer.Flush(); err != nil {
			fmt.Println(err)
		}
	}()
	for item := range errorChannel {
		if _, err := writer.WriteString(item); err != nil {
			fmt.Println(err)
		}
	}
}
