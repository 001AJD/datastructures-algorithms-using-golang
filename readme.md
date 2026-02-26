# Data Structures and Algorithms in Go

Welcome to this repository! This project serves as a practical collection of Go (Golang) implementations, covering fundamental language features, data structures, and algorithmic problems. It is designed to demonstrate key concepts through working code examples.

## Project Structure

The repository is organized into specific modules, each focusing on a distinct topic. For detailed explanations, see the `.md` files maintained within each directory.

- [`/alert-hub`](./alert-hub): Demonstrates the use of interfaces in Go to build a flexible, polymorphic notification system.
- [`/context-package`](./context-package): Explores the Go `context` package for managing timeouts and cancellation in concurrent operations.
- [`/csv-to-json`](./csv-to-json): A CLI tool utilizing the fan-out/fan-in concurrency pattern to efficiently process and convert large CSV files to JSON.
- [`/data-types-loops`](./data-types-loops): Covers fundamental Go syntax, variable declarations, type reflection, and basic control structures.
- [`/goroutines`](./goroutines): An introduction to concurrent programming using goroutines and `sync.WaitGroup` to manage parallel tasks.
- [`/pointers`](./pointers): Illustrates memory management and pointer receivers using a practical "Family Vault" example to modify shared state.
- [`/singly-linked-list`](./singly-linked-list): Implements a custom singly linked list data structure with node insertion and traversal.
- [`/student-grade-manager`](./student-grade-manager): Demonstrates structs, slices, JSON serialization, and file I/O to manage and filter student data.
- [`/url-status-checker`](./url-status-checker): Leverages bounded concurrency to efficiently verify the HTTP status of large lists of URLs in parallel.

## Topics Covered (Completed)

- [x] **Basics:** Variable declarations, data types, `for` loops, `range`, and functions. → [`/data-types-loops`](./data-types-loops)
- [x] **Memory Management:** Pointers (`&`, `*`), memory addresses, and pointer receivers. → [`/pointers`](./pointers)
- [x] **Structs & Data Handling:** Slices, arrays, maps, and JSON serialization. → [`/student-grade-manager`](./student-grade-manager)
- [x] **Interfaces:** Decoupling, polymorphism, and implementing interfaces with structs. → [`/alert-hub`](./alert-hub)
- [x] **Concurrency:** Goroutines, `sync.WaitGroup`, and parallel execution. → [`/goroutines`](./goroutines)
- [x] **Concurrency Patterns:** Fan-out/Fan-in, bounded concurrency, and worker pools. → [`/csv-to-json`](./csv-to-json), [`/url-status-checker`](./url-status-checker)
- [x] **Context:** Timeouts, cancellation, and context propagation. → [`/context-package`](./context-package)
- [x] **File I/O:** CSV parsing, file creation, and output writing. → [`/student-grade-manager`](./student-grade-manager), [`/csv-to-json`](./csv-to-json)
- [x] **Data Structures:** Singly Linked Lists (insertion, traversal). → [`/singly-linked-list`](./singly-linked-list)

## TODO: Learning Roadmap (Meticulous Progress)

### 🟢 1. Intermediate Language Depth

- [ ] **Error Handling:**
  - [ ] Custom Error Types & Wrapping (`errors.Is`, `errors.As`, `%w`)
  - [ ] Deferred Functions (Advanced patterns like named return value modification)
  - [ ] Panic & Recover (Safe recovery in goroutines)
- [ ] **Generics:**
  - [ ] Type parameters and constraints (Go 1.18+)
  - [ ] Creating Generic data structures and functions
- [ ] **Advanced Concurrency:**
  - [ ] `select` with `default` for non-blocking operations
  - [ ] Buffered vs. Unbuffered Channels (and the "Nil Channel" behavior)
  - [ ] Directional Channels (`chan<-` and `<-chan`)
  - [ ] `sync` Package Deep Dive (`sync.Once`, `sync.Pool`, `sync.Cond`, `sync.Map`)
  - [ ] Atomic Operations (`sync/atomic`)
  - [ ] Race Condition Detection & Prevention (`go run -race`)
- [ ] **Reflect & Unsafe:**
  - [ ] Reflection (`reflect` package) basics for generic-like behavior
  - [ ] Understanding `unsafe` package (Pointer arithmetic)

### 🟡 2. Tooling & Ecosystem

- [ ] **Testing & Quality:**
  - [ ] Table-Driven Unit Tests
  - [ ] Mocking Dependencies (using Interfaces or tools like `mockgen`)
  - [ ] Integration Testing (using Test Containers)
  - [ ] Benchmarking (`testing.B`) and Fuzzing (`testing.F`)
  - [ ] Code Coverage Analysis
  - [ ] Profiling (`pprof`) and Trace analysis
- [ ] **Go Modules & Workspace:**
  - [ ] Multi-module Workspaces (`go.work`)
  - [ ] Managing Private Dependencies
  - [ ] Vendoring (`go mod vendor`)

### 🟠 3. Systems & Backend Engineering

- [ ] **Networking & APIs:**
  - [ ] Standard Library `net/http` (Custom Handlers, Middlewares)
  - [ ] RESTful API Design (using frameworks like Chi, Gin, or Fiber)
  - [ ] gRPC & Protocol Buffers
  - [ ] WebSockets and Real-time communication
- [ ] **Data Persistence:**
  - [ ] Database Drivers (`sql` package) vs. ORMs (`GORM`, `ent`)
  - [ ] Database Migrations (e.g., `golang-migrate`)
  - [ ] Connection Pooling & Transaction Management
  - [ ] NoSQL Integration (Redis, MongoDB)
- [ ] **Microservices:**
  - [ ] Message Queues (RabbitMQ, Kafka)
  - [ ] Distributed Tracing (OpenTelemetry)
  - [ ] Dockerization and Kubernetes (K8s) Operators

### 🔵 4. Advanced Algorithms & Data Structures

- [ ] **Linear Structures:**
  - [ ] Doubly Linked Lists
  - [ ] Circular Linked Lists
  - [ ] Stacks & Queues (Implementation using Slices and Linked Lists)
- [ ] **Non-Linear Structures:**
  - [ ] Binary Search Trees (BST) & AVL Trees
  - [ ] Heaps (Min/Max) and Priority Queues
  - [ ] Graphs (Adjacency Matrix/List, BFS, DFS)
  - [ ] Tries (Prefix Trees)
- [ ] **Sorting & Searching:**
  - [ ] Quick Sort, Merge Sort, Heap Sort
  - [ ] Binary Search (and its variants)
- [ ] **Dynamic Programming:**
  - [ ] Memoization & Tabulation patterns in Go

## Prerequisites

- **Go**: Version 1.25.7 or higher.

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
    _Or specifically:_
    ```bash
    go run main.go
    ```
