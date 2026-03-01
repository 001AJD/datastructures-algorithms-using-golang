## Project: Task-Master CLI (Error Handling Lab)

### Objective

Build a command-line tool that processes a list of tasks from a JSON file. The goal is to implement a robust error-handling strategy using **Custom Types**, **Sentinel Errors**, and **Wrapping**.

---

### Problem Statement

You are tasked with creating a "Task Executor" that reads a `tasks.json` file. Each task has an `ID`, a `Name`, and a `Command`. Your program must handle failures gracefully by providing specific context for each type of error encountered during execution.

#### Functional Requirements

1. **File Validation**: If the `tasks.json` file is missing, return a wrapped error providing the file path.
2. **Parsing**: If the JSON is malformed, wrap the standard library error with a custom "Configuration Error" message.
3. **Task Lookup**: If a user requests a specific Task ID that doesn't exist, return a **Sentinel Error** (`ErrTaskNotFound`).
4. **Execution Failure**: Create a **Custom Error Struct** called `TaskExecutionError`. It should store:

- The `TaskID` (int)
- The `ExitCode` (int)
- The original underlying error (wrapped)

---

### Learning Milestones

To successfully complete this project, you must demonstrate the following Go patterns:

- **Sentinel Check**: Use `errors.Is` to check if the error is exactly `ErrTaskNotFound`.
- **Context Wrapping**: Use `fmt.Errorf("... %w", err)` when the file loader fails.
- **Type Assertion**: Use `errors.As` in your `main` function to catch a `TaskExecutionError` and print a specialized message: _"Task 101 failed with exit code 5"_.
- **Interface Implementation**: Ensure your `TaskExecutionError` struct correctly implements the `Error() string` method.

---

### Expected Output Logic

Your `main` function should behave like a router for errors:

- If it's a **Sentinel**, print a "User Warning."
- If it's a **Custom Type**, print the "Technical Metadata" (ID and Exit Code).
- If it's **Unknown**, print a "Generic Failure" message.

---

### Current Progress

- [x] **File Validation**: Custom error type `ErrTaskFileUnavailable` implemented with path and operation context.
- [x] **Custom Error Types**: `ErrInvalidTaskID` implemented for business logic validation.
- [ ] **Parsing**: Malformed JSON needs to be wrapped with a "Configuration Error" message.
- [ ] **Task Lookup**: Sentinel error `ErrTaskNotFound` and lookup logic are pending.
- [ ] **Execution Failure**: `TaskExecutionError` struct needs to be implemented.
- [ ] **Main Function Routing**: Update `main.go` to handle sentinels and the new custom execution error.
