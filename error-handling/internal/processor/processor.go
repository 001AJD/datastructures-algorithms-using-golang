package processor

import (
	"encoding/json"
	"error-handling/internal/domain_errors"
	"error-handling/internal/filesystem"
	"errors"
	"os"
)

type Task struct {
	TaskId   uint8  `json:"taskId"`
	Duration string `json:"duration"`
}

func Init(filepath string) (bool, error) {
	bufferedReader, err := filesystem.NewBufferedFileReader(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			var pathError *os.PathError
			if errors.As(err, &pathError) {
				return false, &domain_errors.ErrTaskFileUnavailable{
					Message:   err.Error(),
					Err:       err,
					Path:      pathError.Path,
					Operation: pathError.Op,
				}
			}

		} else if errors.Is(err, os.ErrPermission) {
			return false, err
		}
	}

	decoder := json.NewDecoder(bufferedReader.File)
	if _, err := decoder.Token(); err != nil {
		return false, err
	}

	// process each json object
	for decoder.More() {
		var t *Task
		if err := decoder.Decode(&t); err != nil {
			return false, err
		}
		if t.TaskId == 20 {
			return false, &domain_errors.ErrInvalidTaskID{
				Err:     domain_errors.ErrReservedTaskID,
				Message: "Task ID is reserved for systems usage.",
			}
		}
	}
	return true, nil
}
