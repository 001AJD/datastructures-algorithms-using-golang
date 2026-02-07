# Pointers & Shared State Management

## Problem Statement
In software development, efficient memory management and state consistency are critical. We face two main challenges:

1.  **Memory Efficiency:** Passing large data structures to functions by value creates copies, which wastes memory and processing power.
2.  **State Mutation:** When representing a real-world entity (like a Bank Account or Wallet), multiple operations (Deposit, Spend) must modify the **original** instance. If we work with copies, updates are lost.

## Solution Explanation

### 1. The Pointer Concept
Go allows us to work directly with memory addresses using **Pointers**.
*   **Address Operator (`&`):** Gets the memory location of a variable.
*   **Dereference Operator (`*`):** Accesses the value stored at a specific memory address.
*   **`new()` function:** Allocates memory for a type and returns a pointer to it.

### 2. The Family Wallet (Shared State)
We implement a `Wallet` struct to demonstrate practical pointer usage.

#### Data Model
*   **Owner:** String (Name of the family).
*   **Balance:** Float32 (Current money available).

#### Logic Flow (Pointer Receivers)
The critical part of this solution is how methods are defined:
```go
func (w *Wallet) Deposit(amount float32) { ... }
```
We use a **Pointer Receiver** (`*Wallet`).
*   **Without Pointer:** The method would receive a *copy* of the wallet. Adding money would only affect the copy.
*   **With Pointer:** The method receives the *address* of the wallet. It modifies the actual `Balance` in memory, ensuring changes persist.

## Visual Flow

### Basic Pointer Mechanics

```text
   Variable 'num'           Pointer 'ptr'
   Address: 0x500           Address: 0x900
  +---------------+        +---------------+
  | Value: 25     | <----- | Value: 0x500  |
  +---------------+        +---------------+
                           (Points to 'num')
```

### Mutating the Wallet

```text
      Initialize Wallet
      Balance: $500.00
             |
             v
   +-------------------+
   |   Family Wallet   | <--- Reference (0xABC)
   +-------------------+
             ^
             |
  +-------------------------+
  | Deposit(200) Called     |
  | Uses pointer 0xABC      |
  | Finds original object   |
  | Updates Balance to $700 |
  +-------------------------+
             |
             v
  +-------------------------+
  | Spend(300) Called       |
  | Uses pointer 0xABC      |
  | Updates Balance to $400 |
  +-------------------------+
```
