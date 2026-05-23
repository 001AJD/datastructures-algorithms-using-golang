# Singly Linked List Implementation

## Problem Statement
Standard arrays or slices in programming store elements in contiguous memory locations. While efficient for access, they have limitations:
1.  **Fixed Size / Reallocation:** Growing an array often requires allocating a new, larger block of memory and copying elements.
2.  **Inefficient Insertions:** Inserting an element at the beginning requires shifting all other elements.

We need a data structure that allows:
*   Dynamic memory allocation (nodes can be anywhere in memory).
*   Efficient insertion at the beginning (O(1) time complexity).
*   Easy traversal of data.

## Solution Explanation

### 1. Data Modeling
We define the structure of our data using two `structs`:

*   **Node:** The fundamental building block.
    *   `Data`: Stores the actual value (type `uint8`).
    *   `Next`: A pointer to the next `Node` in the sequence. If this is `nil`, it signifies the end of the list.
*   **LinkedList:** A wrapper to manage the list.
    *   `Head`: Points to the first node in the list.
    *   `Size`: Tracks the total number of nodes.

### 2. Core Operations

#### Initialization
*   `NewLinkedList()`: Creates an empty list (`Head` is `nil`, `Size` is 0).
*   `NewNode(value)`: Creates a new node instance with the given data and its `Next` pointer initialized to `nil`.

#### Insertion
*   `InsertNodeAtStart(node)`: Adds a new node to the front of the list.
    *   If the list is empty, the new node becomes the Head.
    *   If not empty, the new node points to the current Head, and then becomes the new Head.

#### Traversal
*   `TraverseLinkedList()`: Iterates through the list starting from the `Head`.
    *   It prints the data and memory address of each node.
    *   It follows the `Next` pointer until it reaches `nil`.

## Visual Flow

### Logical Structure
A Singly Linked List connects nodes in one direction.

```text
+------------+       +-------------+       +-------------+
| LinkedList |       |    Node     |       |    Node     |
+------------+       +-------------+       +-------------+
| Head       |------>| Data: 10    |       | Data: 20    |
| Size: 2    |       | Next        |------>| Next        |-----> nil
+------------+       +-------------+       +-------------+
```

### Insertion at Start Flow

```text
   New Node (20)           Current List (Head -> 10)
+-------------+           +-------------+
| Data: 20    |           | Data: 10    |
| Next: nil   |           | Next: nil   |
+-------------+           +-------------+
       |                         ^
       |                         |
       +-------------------------+
       1. New Node points to current Head

       2. Update Head to New Node
          (Head -> 20 -> 10 -> nil)
```
