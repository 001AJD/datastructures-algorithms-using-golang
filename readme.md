# Data Structures and Algorithms in Go

Welcome to this repository! This project serves as a practical collection of Go (Golang) implementations, covering fundamental language features, data structures, and algorithmic problems. It is designed to demonstrate key concepts through working code examples.

## Project Structure

The repository is organized into specific modules, each focusing on a distinct topic:

### 1. Data Types & Loops (`/data-types-loops`)
*   **Concepts:** Basic syntax, variable declarations (`int`, `float`, `string`), Type reflection, and Control structures.
*   **Key Features:**
    *   `for` loops (standard and `while`-style).
    *   Iterating over Slices and Maps using `range`.
    *   Basic conditional logic (Functions).

### 2. Pointers & Memory Management (`/pointers`)
*   **Concepts:** Memory addresses, dereferencing, and modifying shared state.
*   **Key Features:**
    *   **Family Wallet:** A real-world example of using pointer receivers to modify a `Wallet` struct's balance (Deposit/Spend) in place, avoiding copy-by-value issues.
    *   Demonstration of `&` (address) and `*` (value) operators.

### 3. Student Grade Manager (`/student-grade-manager`)
*   **Concepts:** Structs, Slices, Arrays, JSON Serialization, and File I/O.
*   **Key Features:**
    *   Manages a dynamic list of `Student` structs.
    *   Calculates averages from fixed-size subject arrays.
    *   **Filtering Logic:** Identifies "Toppers" (Average >= 90).
    *   **Persistence:** exports the filtered data to a `toppers.json` file.
    *   **Bonus:** A Word Counter utility using `strings` and Maps.

### 4. Singly Linked List (`/singly-linked-list`)
*   **Concepts:** Custom Data Structures, Dynamic Memory, Nodes.
*   **Key Features:**
    *   Implementation of a `LinkedList` and `Node` struct.
    *   **Insertion:** Efficiently adding nodes at the start (O(1)).
    *   **Traversal:** Iterating through the list using pointers until `nil`.

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
