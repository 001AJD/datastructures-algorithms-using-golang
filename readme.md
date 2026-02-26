# Data Structures and Algorithms in Go

Welcome to this repository! This project serves as a practical collection of Go (Golang) implementations, covering fundamental language features, data structures, and algorithmic problems. It is designed to demonstrate key concepts through working code examples.

## Project Structure

The repository is organized into specific modules, each focusing on a distinct topic. For detailed explanations, see the `.md` files maintained within each directory.

*   [`/alert-hub`](./alert-hub): Demonstrates the use of interfaces in Go to build a flexible, polymorphic notification system.
*   [`/context-package`](./context-package): Explores the Go `context` package for managing timeouts and cancellation in concurrent operations.
*   [`/csv-to-json`](./csv-to-json): A CLI tool utilizing the fan-out/fan-in concurrency pattern to efficiently process and convert large CSV files to JSON.
*   [`/data-types-loops`](./data-types-loops): Covers fundamental Go syntax, variable declarations, type reflection, and basic control structures.
*   [`/goroutines`](./goroutines): An introduction to concurrent programming using goroutines and `sync.WaitGroup` to manage parallel tasks.
*   [`/pointers`](./pointers): Illustrates memory management and pointer receivers using a practical "Family Vault" example to modify shared state.
*   [`/singly-linked-list`](./singly-linked-list): Implements a custom singly linked list data structure with node insertion and traversal.
*   [`/student-grade-manager`](./student-grade-manager): Demonstrates structs, slices, JSON serialization, and file I/O to manage and filter student data.
*   [`/url-status-checker`](./url-status-checker): Leverages bounded concurrency to efficiently verify the HTTP status of large lists of URLs in parallel.

## Prerequisites

*   **Go**: Version 1.25.7 or higher.

## How to Run

Each module is a self-contained Go module. To run a specific module:

1.  Navigate to the directory:
    ```bash
    cd student-grade-manager
    ```
2.  Run the main file:
    ```bash
    go run .
    ```
    *Or specifically:*
    ```bash
    go run main.go
    ```
