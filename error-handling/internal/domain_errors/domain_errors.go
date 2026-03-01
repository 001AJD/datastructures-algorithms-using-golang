package domain_errors

import "fmt"

type ErrInvalidTaskID struct {
	Message string
	Err     error
}

func (e *ErrInvalidTaskID) Error() string {
	return fmt.Sprintf("Error: %v | Message: %s", e.Err, e.Message)
}

func (e *ErrInvalidTaskID) Unwrap() error {
	return e.Err
}

// Filesystem related error types
type ErrTaskFileUnavailable struct {
	Message   string
	Err       error
	Path      string
	Operation string
}

func (e *ErrTaskFileUnavailable) Error() string {
	return fmt.Sprintf("op=%s | path=%s | cause=%v", e.Operation, e.Path, e.Err)
}

func (e *ErrTaskFileUnavailable) Unwrap() error {
	return e.Err
}
