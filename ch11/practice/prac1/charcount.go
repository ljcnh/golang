package prac1

import (
	"unicode"
)

// 嫌麻删了一个...

func Charcount(str string) map[rune]int {
	counts := make(map[rune]int)
	invalid := 0
	t := []rune(str)
	for _, r := range t {
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}
		counts[r]++
	}
	return counts
}
