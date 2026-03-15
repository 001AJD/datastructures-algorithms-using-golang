package main

import "fmt"

func main() {
	input := []string{"cats", "acts", "tops", "pots"}
	// output => [["cats","acts"],["tops","pots"]]
	result := groupAnagram(input)
	fmt.Println(result)
}

func groupAnagram(strs []string) [][]string {
	result := [][]string{}
	myMap1 := make(map[[26]int][]string)
	// iterate each word in array
	// iterate each char in word
	// store the [26] array as key in hashmap
	// if same key found, then concat the value

	for _, word := range strs {
		frequency := [26]int{}
		for _, alpha := range word {

			// ASCII value substraction, 99-ASCII(c) => 99-97 = 2 => increment value at index 2 in frequency array
			// This will provide [1 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 1 0 0 0 0 0 0] array like this for each word, for anagram this array will be same and can be used as key in hashmap
			// in int array by default all the 26 index has 0 value, increment the value by 1
			frequency[alpha-'a']++
		}
		fmt.Println(frequency)
		value, _ := myMap1[frequency]
		myMap1[frequency] = append(value, word) // simply add the key, value into the hashmap
		fmt.Println(myMap1)
	}

	for _, v := range myMap1 {
		result = append(result, v)
	}

	return result
}
