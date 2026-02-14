package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var client = &http.Client{
	Timeout: 2 * time.Second,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

// CheckStatus initiates a synchronous HTTP GET request for the provided URL.
// It logs the request initiation and the resulting status code (or error) to the console.
//
// Parameters:
//   - url: A string representing the target endpoint (e.g., "https://example.com").
//
// Returns:
//   - A string message detailing whether the request succeeded with a status code
//     or failed with a specific error message.
//
// Note: This function does not currently close the response body, which may lead
// to resource leaks in long-running applications.
func CheckStatus(url string) string {
	fmt.Println("GET request initiated inside go routine for ", url)
	response, err := client.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		return "Request with " + url + " failed with error :: " + err.Error()
	} else {
		defer response.Body.Close()
		fmt.Println("Http GET request successful, statusCode :: ", response.StatusCode)
		return "Request with " + url + " succeed with status code :: " + strconv.Itoa(response.StatusCode)
	}
}
