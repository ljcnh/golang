package main

import (
	"fmt"
)

func main() {
	s := "asdasdsadasd"
	fmt.Println(comma(s))
	a := []byte(s)
	fmt.Println(a)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
