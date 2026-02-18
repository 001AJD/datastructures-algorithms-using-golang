package main

import "fmt"

func main() {
	fmt.Println("CSV to json transformation started...")
	var path string = "./data/top-100.csv"
	if ProcessFile(path) {
		fmt.Println("File processing complete")
	} else {
		fmt.Println("File processing failed")
	}
}
