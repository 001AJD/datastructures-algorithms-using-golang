package main

import "fmt"

// This function should take the integer item to be searched as an input, and a searchSpace array of integers
// returns boolean, true if the item is present in the list, false if the items is not present
func BinarySearch(item int, searchSpace []int) bool {

	var left, right, mid int
	left = 0
	right = len(searchSpace) - 1
	for {
		mid = left + (right-left)/2
		fmt.Println("mid = ", mid)
		fmt.Println("left = ", left)
		fmt.Println("right = ", right)
		fmt.Println("******************")
		if left > right {
			return false
		}
		if searchSpace[mid] == item {
			return true
		}
		if searchSpace[mid] > item {
			right = mid - 1
		}
		if searchSpace[mid] < item {
			left = mid + 1
		}
	}

}
