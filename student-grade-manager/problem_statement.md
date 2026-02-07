# Student Grade Manager & Word Counter

## Problem Statement
We need to build a system to manage student grades for a classroom. The specific requirements are:

1.  **Dynamic Class Size:** The number of students can change (grow or shrink), so we cannot use a fixed-size array for the student list.
2.  **Fixed Subjects:** Each student has scores for exactly **3 subjects**.
3.  **Performance Tracking:** We need to calculate the average score for each student.
4.  **Identify Toppers:** Filter out students who have an average score of **90 or higher**.
5.  **Data Persistence:** Save the list of toppers to a file (`toppers.json`) for record-keeping.
6.  **Word Analysis:** (Bonus) A utility to count the frequency of words in a text string.

## Solution Explanation

### 1. Data Modeling
We use a Go `struct` to define a **Student**. This keeps related data together.
*   **Name:** String
*   **Subjects:** A fixed array `[3]float32` (since the number of subjects is constant).
*   **Average:** Float32 (calculated later).

### 2. Managing the List
Since the number of students varies, we use a Go **Slice** (`[]Student`). Slices are dynamic arrays that can grow as we `append` new students.

### 3. Logic Flow
The program follows these steps:
1.  **Initialize:** Create a slice of students with hardcoded data.
2.  **Process:** Iterate through each student:
    *   Sum the 3 subject scores.
    *   Divide by 3 to get the average.
    *   Store the average in the student struct.
3.  **Filter:** If `Average >= 90`, add the student to a separate `toppers` slice.
4.  **Export:** Use the `encoding/json` package to convert the `toppers` slice into JSON text and write it to `toppers.json`.

### 4. Word Counter
A separate function `CountWord` takes a string, splits it into words, and uses a `map[string]int` to count how many times each word appears.

## Visual Flow

```text
Start: List of Students
       |
       v
+-------------+
| Loop through|<---------------------------+
|  Students   |<-------------------+       |
+-------------+                    |       |
       |                           |       |
       v                           |       |
+-----------------------+          |       |
| Calculate Average     |          |       |
|       Score           |          |       |
+-----------------------+          |       |
       |                           |       |
       v                           |       |
+-----------------------+    No    |       |
|    Average >= 90?     |----------+       |
+-----------------------+                  |
       | Yes                               |
       v                                   |
+-----------------------+                  |
|  Add to Toppers List  |------------------+
+-----------------------+
       |
       v (End of List)
+-----------------------+
| Convert Toppers to    |
|       JSON            |
+-----------------------+
       |
       v
+-----------------------+
| Write to toppers.json |
+-----------------------+
```