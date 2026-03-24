package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 4, 100, 200, 300, 400}
	var hmap map[int]struct{} = make(map[int]struct{})

	for i := range arr {
		hmap[arr[i]] = struct{}{}
	}

	var maxLength int = 1
	for i := 0; i < len(arr); i++ {
		// check if the current number - 1 exists in the set, if not it is start of a sequence
		_, ok := hmap[arr[i]-1]
		if !ok {
			currentNum := arr[i]
			currentLength := 1
			for {
				_, ok := hmap[currentNum+currentLength]
				if ok {
					currentLength++
				} else {
					break
				}
			}
			maxLength = max(maxLength, currentLength)
		}
	}
	fmt.Println(maxLength)
}
