package main

import "fmt"

func main() {
	// var slice []int = []int{10, 20, 30, 40, 50, 60}
	var slice2 []int = []int{60}
	var item int = 60
	result := BinarySearch(item, slice2)
	fmt.Println(result)
}
