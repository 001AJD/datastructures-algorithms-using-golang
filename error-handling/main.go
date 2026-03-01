package main

import (
	"error-handling/internal/domain_errors"
	"error-handling/internal/processor"
	"errors"
	"fmt"
)

func main() {

	fmt.Println("starting task handler...")
	// result, err := processor.Init("./sample-data/tasks.json")
	// result, err := processor.Init("./sample-data/invalid_structure.json")
	result, err := processor.Init("./sample-data/tasks.json")
	if err != nil {
		// filesystem error
		var errTaskFileUnavailable *domain_errors.ErrTaskFileUnavailable

		if errors.As(err, &errTaskFileUnavailable) {
			fmt.Println("Error:: File does not exist")
		} else if errors.Is(err, domain_errors.ErrReservedTaskID) {
			fmt.Println("Error: Reserved Task ID found in the data")
		} else {
			fmt.Println("Error occurred")
			fmt.Println(err)
		}
	} else {
		fmt.Println("File processing completed successfully:: ", result)
	}
}
