package processor

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
func CSVToJSON(record []string) (Domains, error) {
	var result Domains
	i, err := strconv.Atoi(record[1])
	if err != nil {
		fmt.Println("string conversion failed!!")
		return Domains{}, fmt.Errorf("Bad record, failed string conversion %s", err.Error())
	}
	result.Domain = record[0]
	result.Popularity = int16(i)
	return result, nil
}
