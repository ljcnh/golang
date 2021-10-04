package main

import "fmt"

func main() {
	fmt.Println(compare("sd", "sad"))
	fmt.Println(compare("asd", "sad"))
}

func compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range s1 {
		m1[v]++
	}
	for _, v := range s2 {
		m2[v]++
	}
	for i, v := range m1 {
		if m2[i] != v {
			return false
		}
	}
	return true
}
