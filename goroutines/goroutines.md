# Go Routines Module Summary

This document summarizes the functionality of the `go-routines` module, which demonstrates concurrent programming using Go routines and `sync.WaitGroup`.

## `go.mod`

The `go.mod` file specifies the module name as `go-routines` and indicates that it uses Go version `1.25.7`.

## `main.go`

The `main.go` file implements a basic order management system utilizing Go routines for concurrent processing.

### Structures

- **`Orders`**: A struct representing an order with the following fields:
    - `ID`: A `uint8` representing the unique identifier of the order.
    - `Status`: A `string` indicating the current status of the order (e.g., "Pending", "Processing", "Shipped", "Delivered").

### Core Functionality

The `main` function orchestrates the concurrent operations:

1.  **Initialization**:
    -   Prints "Go Routines".
    -   Initializes a `sync.WaitGroup` to synchronize the execution of goroutines. Two goroutines are expected, so `wg.Add(2)` is called.
    -   Calls `createOrders(10)` to generate 10 initial orders, all with "Pending" status. This function simulates order creation with random delays.

2.  **Goroutines Launch**:
    -   **`printOrders` Goroutine**: Launched concurrently to iterate through the list of orders and print their current status. It includes comments explaining Go's automatic pointer dereferencing when accessing struct fields.
    -   **`updateOrders` Goroutine**: Launched concurrently to randomly update the status of each order to "Processing", "Shipped", or "Delivered" after a simulated delay.

3.  **Synchronization**:
    -   `wg.Wait()` is called to block the main function until both `printOrders` and `updateOrders` goroutines have completed their execution.
    -   Once all operations are finished, it prints "All the operations are completed!".

### Flow Diagram

```text
+---------------------+
| main function start |
+---------------------+
          |
          v
+-------------------------+
| Initialize WaitGroup    |
| and Add(2)              |
+-------------------------+
          |
          v
+-------------------------+
| createOrders(10)        |
+-------------------------+
          |
          +-----------------+
          |                 |
          v                 v
+-------------------+ +-------------------+
| goroutine:        | | goroutine:        |
| printOrders       | | updateOrders      |
+-------------------+ +-------------------+
          |                 |
          v                 v
+-------------------+ +-------------------+
| defer wg.Done()   | | defer wg.Done()   |
+-------------------+ +-------------------+
          |                 |
          +--------+--------+
                   |
                   v
          +------------------+
          |    wg.Wait()     |
          +------------------+
                   |
                   v
+------------------------------------+
| All the operations are completed!  |
+------------------------------------+
                   |
                   v
          +------------------+
          | main function end|
          +------------------+
```

### Helper Functions

-   **`createOrders(count uint8) []*Orders`**:
    -   Takes a `count` as input.
    -   Creates a slice of `Orders` pointers.
    -   Simulates creating `count` orders, assigning them an `ID` and an initial `Status` of "Pending".
    -   Introduces random `time.Sleep` to simulate work.
    -   Prints the creation of each order.
    -   Returns the slice of newly created orders.

-   **`printOrders(orders []*Orders)`**:
    -   Takes a slice of `Orders` pointers.
    -   Iterates through each order and prints its index and status.
    -   Demonstrates how Go automatically dereferences pointers when accessing fields (`v.Status`) and also shows explicit dereferencing (`(*v).Status`).
    -   Includes a random `time.Sleep` to simulate processing time.

-   **`updateOrders(orders []*Orders)`**:
    -   Takes a slice of `Orders` pointers.
    -   Iterates through each order.
    -   Randomly assigns a new status ("Processing", "Shipped", or "Delivered") to the order after a simulated delay.
    -   Prints the update of each order's status.

This module provides a clear, practical example of how to use Go routines for concurrent tasks and `sync.WaitGroup` for managing their completion.