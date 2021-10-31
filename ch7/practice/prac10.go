package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(sort.StringSlice{"a", "b", "a", "d", "z", "a", "c", "d", "a", "a"}))
	s := "pqqpasdgd"
	fmt.Println(IsPalindrome(to(s)))
	s = "aaadddaaa"
	fmt.Println(IsPalindrome(to(s)))
}

func to(s string) sort.Interface {
	var f []string
	for i := 0; i < len(s); i++ {
		f = append(f, string(s[i]))
	}
	return sort.StringSlice(f)
}
