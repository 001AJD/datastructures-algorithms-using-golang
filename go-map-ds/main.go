package main

import "fmt"

func main() {
	fmt.Println(isAnagram("jar", "jam"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]uint32)
	m2 := make(map[byte]uint32)
	for i := 0; i < len(s); i++ {
		m1[s[i]]++
		m2[t[i]]++
		// value, ok := m1[s[i]]
		// if ok {
		// 	m1[s[i]] = value + 1
		// }
		// value2, ok := m2[t[i]]
		// if ok {
		// 	m2[t[i]] = value2 + 1
		// }
	}
	fmt.Println(m1)
	fmt.Println(m2)

	for k, _ := range m1 {
		if m2[k] != m1[k] {
			return false
		}
	}
	return true
}
