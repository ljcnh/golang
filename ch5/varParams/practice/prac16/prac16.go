package prac16

import "fmt"

func main() {
	s := []string{"4324", "dgrp[", "dfdk[wep"}
	fmt.Println(MyJoin("...", "324", "324", "sdafs", "?><<"))
	fmt.Println(MyJoin("...", s...))
}

func MyJoin(sep string, str ...string) string {
	var res string
	for i, s := range str {
		res += s
		if i != len(str)-1 {
			res += sep
		}
	}
	return res
}
