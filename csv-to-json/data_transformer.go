package main

import (
	"fmt"
	"strconv"
)

type Domains struct {
	Domain     string `json:"domainName"`
	Popularity int16  `json:"popularity"`
}

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
