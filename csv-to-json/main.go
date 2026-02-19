package main

import (
	"csv-to-json/internal/processor"
	"fmt"
)

func main() {
	fmt.Println("CSV to json transformation started...")
	var inputFilePath string = "./data/top-100.csv"
	var outputFilePath string = "./output/output.json"
	if _, err := processor.ProcessFile(inputFilePath, outputFilePath); err != nil {
		fmt.Println(fmt.Errorf("File processing failed with error %w", err))
	} else {
		fmt.Println("File processing completed with success")
	}
}
