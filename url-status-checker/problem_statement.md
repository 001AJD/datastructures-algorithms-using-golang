The `url-status-checker` is a command-line tool designed to concurrently check the status of a list of URLs.

The core problem this project addresses is how to efficiently check the status of a large number of URLs from a given file without checking them one by one sequentially. The key is to leverage Go's concurrency features (goroutines) to perform these checks in parallel, which significantly speeds up the process.

The program works as follows:
1. It reads a list of domains from a CSV file.
2. It divides the list of domains into smaller batches.
3. For each batch, it spawns a new goroutine to process it.
4. Within each batch-processing goroutine, it further spawns a separate goroutine for each individual domain to check its HTTP status.
5. It uses `net/http` to send a GET request to each domain and determines if it's reachable based on the HTTP response.

The primary purpose is to demonstrate the use of goroutines to parallelize I/O-bound tasks, such as making network requests, to improve performance.

## Architecture Diagram

```
+-----------------+
|   main func     |
+-----------------+
        |
        v
+-----------------+
|  ProcessFile    |
| (top-100.csv)   |
+-----------------+
        |
        | Divides into batches of 20
        |
        v
+-------------------------------------------------+
|                                                 |
| +-----------------+ +-----------------+ +-----+ |
| |  Batch 1 (20)   | |  Batch 2 (20)   | | ... | |
| | (Goroutine)     | | (Goroutine)     | | ... | |
| +-----------------+ +-----------------+ +-----+ |
|                                                 |
+-------------------------------------------------+
        |
        | For each batch, spawn goroutines for each domain
        v
+-------------------------------------------------+
|                                                 |
| +-----------------+ +-----------------+ +-----+ |
| | Domain 1 (GR)   | | Domain 2 (GR)   | | ... | |
| | CheckStatus()   | | CheckStatus()   | | ... | |
| +-----------------+ +-----------------+ +-----+ |
|                                                 |
+-------------------------------------------------+
        |
        v
+-----------------+
|   HTTP GET      |
+-----------------+
        |
        v
+-----------------+
|  Print Status   |
+-----------------+
```
