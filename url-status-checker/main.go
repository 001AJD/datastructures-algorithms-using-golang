package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	fmt.Println("URL status checker")

	var urls []string = []string{
		"https://www.google.com",
		"https://www.githubb.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.x.com", // twitter
		"https://linkedin.com",
		"https://yahoo.com",
	}
	var results []string

	wg.Add(len(urls))
	start1 := time.Now() // capture time when goroutines are used, parallel processing
	for i, v := range urls {
		fmt.Println(i, v)

		go func(url string) {
			defer wg.Done()
			r := checkStatus(url)

			mu.Lock()
			results = append(results, r)
			mu.Unlock()

		}(v) // notice that v is being passed and FUNCTION PARAMETER to the goroutine
	}
	wg.Wait()
	elapsed1 := time.Since(start1)

	// sequential status check without goroutine
	start2 := time.Now()
	for _, v := range urls {
		checkStatus(v)
	}
	elapsed2 := time.Since(start2)

	for i, v := range results {
		fmt.Printf("%d ==> %s\n", i, v)
	}
	fmt.Println()
	fmt.Println("Parallel :: Status checker took " + elapsed1.String() + " to complete")
	fmt.Println("Sequential :: Status checker took " + elapsed2.String() + " to complete")
}

func checkStatus(url string) string {
	fmt.Println("GET request initiated inside go routine for ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return "Request with " + url + " failed with error :: " + err.Error()
	} else {
		fmt.Println("Http GET request successful, statusCode :: ", response.StatusCode)
		return "Request with " + url + " succeed with status code :: " + strconv.Itoa(response.StatusCode)
	}
}
