package main

import (
	"fmt"
	"unicode"
)

func main() {
	var b []byte = []byte{24, ' ', 46, 21, ' ', ' ', ' ', 4}
	b = removeMul1(b)
	fmt.Println(b)
	b = []byte{24, ' ', 46, 21, ' ', ' ', ' ', 4}
	b = removeMul2(b)
	fmt.Println(b)
}

func removeMul1(b []byte) []byte {
	if len(b) < 2 {
		return b
	}
	var i = 0
	for _, v := range b {
		if v == byte(' ') && i > 0 && b[i-1] == byte(' ') {
			continue
		}
		b[i] = v
		i++
	}
	return b[:i]
}

func removeMul2(b []byte) []byte {
	if len(b) < 2 {
		return b
	}
	var end = 0
	for j := 1; j < len(b); j++ {
		if unicode.IsSpace(rune(b[j])) && unicode.IsSpace(rune(b[end])) {
			continue
		}
		end++
		b[end] = b[j]
	}
	return b[:end+1]
}
