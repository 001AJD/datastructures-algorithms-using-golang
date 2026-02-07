package main

import "strings"

func CountWord(str string) map[string]int {
	var report map[string]int = map[string]int{}
	for word := range strings.FieldsSeq(str) {
		report[string(word)]++
	}
	return report
}
