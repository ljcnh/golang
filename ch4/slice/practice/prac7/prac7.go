package main

import "fmt"

func main() {
	var b []byte = []byte{24, 'v', 46, 21, 43, 5, 0, 4}
	b = reverse(b)
	fmt.Println(b)
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
