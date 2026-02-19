package processor

import "sync"

// worker consumes CSV records from jobs, converts each record to a Domains
// object, and sends the transformed value to processedRows.
func worker(jobs <-chan []string, processedRows chan<- Domains, errorChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range jobs {
		if outputJson, err := CSVToJSON(record); err != nil {
			// put error message in the error channel
			errorChannel <- err.Error()
		} else {
			processedRows <- outputJson
		}
	}
}
