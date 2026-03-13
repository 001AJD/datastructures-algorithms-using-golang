package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := []string{"", "World"}
	encodedString := encode(input)
	fmt.Printf("Encoded String :: %s\n", encodedString)

	decodedArr := decode(encodedString)
	fmt.Printf("Decoded array :: %s\n", decodedArr)

	res := len("")
	fmt.Println(res)

}

func encode(str []string) string {
	var result string
	for _, v := range str {
		result = result + strconv.Itoa(len(v)) + "#" + v
	}
	return result
}

func decode(encoded string) []string {
	var result []string
	var wordLengthStr string
	var wordLengthInt int = 0
	for i := 0; i < len(encoded); i++ {

		if encoded[i] == '#' {

			wordLengthInt, _ = strconv.Atoi(wordLengthStr)
			wordLengthStr = ""

			// iterate wordLengthInt time from here onwards
			// fmt.Println("Wordlength :: ", wordLengthStr)
			if wordLengthInt == 0 {
				result = append(result, "")
			} else {
				i++
				var word string
				for j := 0; j < wordLengthInt; j++ {
					word = word + string(encoded[i])
					i++
				}
				fmt.Println(word)
				result = append(result, word)
				i--
			}

		} else {
			wordLengthStr = wordLengthStr + string(encoded[i])
		}
	}
	return result
}
