package main

import "fmt"

func main() {
	result := scoreOfString("neetcode")
	fmt.Println(result)
}

func scoreOfString(s string) int {
	var result int = 0
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = diff * -1
		}
		fmt.Printf("\n%d - %d = %d\n", s[i], s[i-1], diff)
		result += diff
	}
	return result
}
