package logging

import (
	"csv-to-json/internal/filesystem"
	"fmt"
)

const logFilePath = "./logs/log.txt"

var logFile *filesystem.BufferedFileWriter

func ProcessLogs(errorChannel <-chan string) {

	defer logFile.Close()
	for item := range errorChannel {
		if _, err := logFile.BufferedWriter.WriteString(item); err != nil {
			fmt.Println(err)
		}
	}
}

func InitializeLogger() error {
	file, err := filesystem.NewBufferedFileWriter(logFilePath)
	if err != nil {
		return err
	} else {
		logFile = file
		return nil
	}
}
