# Error in golang

- Errors are values in golang
- Any struct implementing `Error() string` function becomes Error type.

```go
type MyError struct{
  Query string
  Err error
}
func (m *MyError) Error() string {
  return e.Query+": " + e.Err.Error()
}
```

- In above code snippet the struct implements the `Error() string {}` function, hence the struct becomes of type error. This struct can be passed to a function expecting an error type value.

## Implementation Progress
- **Custom Error Types**: Created `ErrInvalidTaskID` and `ErrTaskFileUnavailable` to provide context-rich error reporting.
- **Sentinel Errors**: Implemented `ErrReservedTaskID` as a package-level sentinel error for specific business logic violations.
- **Error Checking**:
    - Utilized `errors.As` for retrieving structured data from wrapped custom errors.
    - Utilized `errors.Is` to check for specific sentinel errors (`ErrReservedTaskID`) and standard library errors (`os.ErrNotExist`).
- **Error Wrapping**: Successfully implemented `Unwrap() error` methods to support error chaining and preservation of the original cause.
