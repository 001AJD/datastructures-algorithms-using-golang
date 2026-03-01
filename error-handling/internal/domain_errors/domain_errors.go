package domain_errors

import "fmt"

// Json data validation related error types
type ErrInvalidTaskID struct {
	Message string
	Err     error
}

func (e *ErrInvalidTaskID) Error() string {
	return "Error message: " + e.Err.Error()
}

type ErrInvalidToken struct {
	Message string
	Err     error
}

func (e *ErrInvalidToken) Error() string {
	return fmt.Sprintf("\nErr: %v, \nMessage: %s", e.Err, e.Message)
}

func (e *ErrInvalidToken) Unwrap() error {
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
	return fmt.Sprintf("\nop=%s \npath=%s \ncause=%v", e.Operation, e.Path, e.Err)
}

func (e *ErrTaskFileUnavailable) Unwrap() error {
	return e.Err
}

type ErrTaskFilePermission struct {
	Message string
	Err     error
}

func (e *ErrTaskFilePermission) Error() string {
	return fmt.Sprintf("\nMessage:%s \nError:%v", e.Message, e.Err)
}

func (e *ErrTaskFilePermission) Unwrap() error {
	return e.Err
}
