// For a given array calculate
// prefix sum
// suffix sum
// Prefix product
// suffix product

package main

import "fmt"

func main() {
	var arr []uint8 = []uint8{1, 2, 4, 6}

	prefixP := PrefixProduct(arr)
	fmt.Printf("Prefix product :: %v\n", prefixP)

	suffixP := SuffixProduct(arr)
	fmt.Printf("Suffix product :: %v\n", suffixP)

	productArr := ArrayProduct(prefixP, suffixP, arr)
	fmt.Printf("Array of product except self :: %v\n", productArr)
}

func PrefixProduct(arr []uint8) []uint8 {
	var result []uint8 = make([]uint8, len(arr)+1)
	result[0] = 1
	for i := 1; i <= len(arr); i++ {

		result[i] = arr[i-1] * result[i-1]

	}
	return result
}

func SuffixProduct(arr []uint8) []uint8 {
	var result []uint8 = make([]uint8, len(arr)+1)
	result[len(result)-1] = 1
	for i := len(arr) - 1; i >= 0; i-- {

		result[i] = arr[i] * result[i+1]

	}
	return result
}

func ArrayProduct(p []uint8, s []uint8, input []uint8) []uint8 {
	var result []uint8 = make([]uint8, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = p[i] * s[i+1]
	}
	return result
}
