The `url-status-checker` is a command-line tool designed to concurrently check the status of a list of URLs.

The core problem this project addresses is how to efficiently check the status of a large number of URLs from a given file without checking them one by one sequentially. The key is to leverage Go's concurrency features (goroutines) to perform these checks in parallel, which significantly speeds up the process.

The program works as follows:
1. It reads a list of domains from a CSV file.
2. It divides the list of domains into smaller batches.
3. For each batch, it spawns a new goroutine to process it.
4. Within each batch-processing goroutine, it further spawns a separate goroutine for each individual domain to check its HTTP status.
5. It uses `net/http` to send a HEAD request to each domain and determines if it's reachable based on the HTTP response.

The primary purpose is to demonstrate the use of goroutines to parallelize I/O-bound tasks, such as making network requests, to improve performance.

## Result of processing 100K domains

```
URL status checker initiating...

Alloc = 0 MB
TotalAlloc = 0 MB
Sys = 8 MB
NumGC = 0
Goroutines: 1
------
2026/02/14 23:54:09 protocol error: received DATA on a HEAD request
2026/02/14 23:55:03 Transport: unhandled response frame type *http.http2UnknownFrame

File processing complete
Processing completed in 3m23.709959333s

Alloc = 848 MB
TotalAlloc = 2586 MB
Sys = 2755 MB
NumGC = 16
Goroutines: 174132
------
```

## Architecture Diagram

```
+-----------------+
|   main func     |
+-----------------+
        |
        v
+-----------------+
|  ProcessFile    |
| (reads domains) |
+-----------------+
        |
        | Sends domains to a channel
        v
+-----------------+
| jobs (buffered  |
|   channel)      |
+-----------------+
        |
        | Workers read from the channel (Bounded Concurrency)
        v
+-------------------------------------------------+
|                                                 |
| +---------------+ +---------------+ +---------+ |
| | Worker 1 (GR) | | Worker 2 (GR) | | ... (N) | |
| | CheckStatus() | | CheckStatus() | | ...     | |
| +---------------+ +---------------+ +---------+ |
|                                                 |
+-------------------------------------------------+
        |
        | For each domain
        v
+-----------------+
|   HTTP HEAD     |
+-----------------+
        |
        v
+-----------------+
|  Print Status   |
+-----------------+
```
