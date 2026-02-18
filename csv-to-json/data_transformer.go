package main

import (
	"fmt"
	"strconv"
)

type Domains struct {
	Domain     string `json:"domainName"`
	Popularity int16  `json:"popularity"`
}

// CSVToJSON converts a CSV record into a Domains struct.
// It expects record[0] to contain the domain name and record[1] to contain
// the popularity value as a numeric string.
func CSVToJSON(record []string) Domains {
	var result Domains
	i, err := strconv.Atoi(record[1])
	fmt.Println(i)
	if err != nil {
		fmt.Println("string converstion failed!!")
	}
	result.Domain = record[0]
	result.Popularity = int16(i)
	fmt.Println(result)
	return result
}
