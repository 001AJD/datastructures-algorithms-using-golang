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
