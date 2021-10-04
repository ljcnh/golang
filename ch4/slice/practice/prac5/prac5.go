package main

import "fmt"

func main() {
	var s []string = []string{"123", "123", "sds", "sds", "sprty", "sprty", "sprty", "sprthrt"}
	s = removeMul(s)
	fmt.Println(s, len(s), cap(s))
}

func removeMul(s []string) []string {
	var i = 0
	for _, v := range s {
		if v != s[i] {
			i++
			s[i] = v
		}
		fmt.Println(i, s[i])
	}
	return s[:i+1]
}
