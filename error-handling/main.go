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
	result, err := processor.Init("./sample-data/bad_data.json")
	if err != nil {
		// filesystem error
		var errTaskFileUnavailable *domain_errors.ErrTaskFileUnavailable
		var errTaskFilePermission *domain_errors.ErrTaskFilePermission
		var errInvalidToken *domain_errors.ErrInvalidToken
		// json data file errors
		var ErrInvalidTaskID *domain_errors.ErrInvalidTaskID

		if errors.As(err, &errTaskFileUnavailable) {
			fmt.Println("Error:: File does not exist")
		} else if errors.As(err, &errTaskFilePermission) {
			fmt.Println("Error:: Permission denined on file")
		} else if errors.As(err, &errInvalidToken) {
			fmt.Println("Error occurred:: Invalid token found in json file")
		} else if errors.As(err, &ErrInvalidTaskID) {
			fmt.Println("Error: Invalid Task ID found in the data")
		} else {
			fmt.Println("Error occurred")
			fmt.Println(err)
		}
	} else {
		fmt.Println("File processing completed successfully:: ", result)
	}

	// for decoder.More() {
	// 	var task Tasks

	// 	err := decoder.Decode(&task)
	// 	if err != nil {
	// 		var typeErr *json.UnmarshalTypeError
	// 		if errors.As(err, &typeErr) {
	// 			if typeErr.Field == "taskId" {

	// 			}
	// 		}
	// 	}
	// 	fmt.Println(task.TaskId, task.Duration)
	// }

}
