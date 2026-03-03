# PRD: Go-Proc-Chain (Interactive Process Monitor)

## 1. Purpose

The goal of this project is to build an interactive Command Line Interface (CLI) that simulates a system process monitor using a **Singly Linked List**. It focuses on mastering pointer manipulation, node deletion, and list traversal in Go.

## 2. Core Features

### 2.1 Data Loading (The "Genesis")

- **Source:** The app must load process data from a local `processes.csv` file.
- **Mechanism:** On startup, the app parses the CSV and performs a "Bulk Append" to build the initial linked list.
- **Node Structure:** Each node represents a Process containing:
  - `PID` (int)
  - `Name` (string)
  - `CPU` (float64)
  - `Memory` (string)
  - `Next` (Pointer to the next Node)

### 2.2 Visualization (The "Monitor")

- **Command:** `list`
- **Requirement:** Traverse the list from Head to Tail and print an ASCII representation of the chain.
- **Format:** `[PID: 101 | system_d] -> [PID: 402 | chrome] -> [PID: 505 | go] -> NULL`

### 2.3 Process Termination (The "Kill")

- **Command:** `kill <PID>`
- **Requirement:** Locate the node with the matching PID and remove it from the list.
- **Logic:** Implement the "bypass" logic where the `PreviousNode.Next` is updated to `TargetNode.Next`.
- **Edge Case:** Correctly handle deleting the Head node (updating the list's head pointer) and the Tail node.

### 2.4 Priority Boosting (The "Boost")

- **Command:** `boost <PID>`
- **Requirement:** Move a specific process to the very top (Head) of the list.
- **Logic:**
  1. Search for the node.
  2. Unlink it from its current position (repairing the gap).
  3. Set its `Next` to the current Head.
  4. Update the List Manager's Head pointer to this node.

## 3. Technical Requirements

### 3.1 Language & Tools

- **Language:** Go (Golang)
- **Packages:** `encoding/csv`, `os`, `fmt`, `bufio`.
- **Concurrency (Optional):** Use a goroutine to simulate "real-time" CPU usage updates by traversing and modifying node data every 5 seconds.

### 3.2 Constraints

- **No Slice/Array Storage:** Once loaded, the data must reside only in the Linked List. No backing arrays are allowed for the core logic.
- **Manual Traversal:** All searches must be performed via `for` loops traversing pointers (`current = current.Next`).

## 4. User Interaction Flow

1. **Launch:** User runs `go run main.go`.
2. **Initialization:** App reads `processes.csv` and displays `"Loaded 5 processes."`
3. **Command Loop:**
   - User types `list` to see the current state.
   - User types `kill 402` to remove Chrome.
   - User types `boost 505` to move the Go engine to the top.
4. **Update:** App re-renders the chain after every successful operation.

## 5. Success Criteria

- The application handles "Kill" operations without losing the rest of the list (no "orphaned" nodes).
- The application handles "Boost" operations for nodes at the head, middle, and tail without crashing.
- The code is modular, with the Linked List logic separated from the CLI/Parser logic.
