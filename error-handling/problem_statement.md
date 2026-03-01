# Task-Master CLI: Project Summary & Technical Notes

## Project Overview

The **Task-Master CLI** is a Go-based utility designed to process task data from JSON files while demonstrating robust error-handling patterns. The project emphasizes the use of custom error types, error wrapping, and the `errors` package functionality (`Is`, `As`, `Unwrap`).

---

## Codebase Summary

### 1. Domain Errors (`internal/domain_errors`)

Defines specialized error structures to provide rich context when failures occur:

- **`ErrInvalidTaskID`**: Used when a task violates business logic (e.g., using a reserved ID).
- **`ErrTaskFileUnavailable`**: Wraps filesystem errors with metadata like `Path` and `Operation`.
- **Implementation**: Both types implement the `Error()` and `Unwrap()` methods, supporting Go's standard error-chaining protocols.

### 2. Filesystem Utility (`internal/filesystem`)

- Provides a `BufferedReader` abstraction to handle file I/O efficiently using `bufio`.
- Ensures files are opened with appropriate permissions and provides a clean interface for closing resources.

### 3. Task Processor (`internal/processor`)

- **Initialization**: Validates the existence of input files and maps low-level `os` errors to domain-specific types.
- **Streaming JSON**: Uses `json.NewDecoder` to process tasks one by one, ensuring low memory overhead for large datasets.
- **Validation**: Implements logic to catch "Reserved Task IDs" during the decoding phase.

### 4. Application Entry Point (`main.go`)

- Acts as the error router.
- Demonstrates **Type Assertion** using `errors.As` to distinguish between filesystem failures and data validation errors, providing tailored feedback to the user.

---

## Technical Notes

- **Error Wrapping**: The project correctly uses wrapping to preserve original context (e.g., `os.PathError` is preserved within `ErrTaskFileUnavailable`).
- **Sentinel Errors**: While the structure is ready, specific sentinel errors (like `ErrTaskNotFound`) are identified as future integration points for specific lookup logic.
- **Data Integrity**: The processor currently flags `TaskId: 20` as a reserved system ID, demonstrating how custom errors can enforce business rules during parsing.
- **Extensibility**: The architecture allows for easy addition of a `TaskExecutionError` struct to handle process-level failures (exit codes, command output) in future iterations.
