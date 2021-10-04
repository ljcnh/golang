package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "-3573246843.523434"
	fmt.Println(comma(s))
	fmt.Println(len(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	sarr := strings.Split(s, ".")
	l := sarr[0]
	for i := 0; i < len(l); i++ {
		buf.WriteByte(l[i])
		if i != n-1 && (n-i-1)%3 == 0 {
			buf.WriteByte(',')
		}
	}
	// 小数也做了同样的操作 不过没必要
	if len(sarr) > 1 {
		buf.WriteByte('.')
		r := sarr[1]
		rn := len(r)
		for i := 0; i < rn; i++ {
			buf.WriteByte(r[i])
			if i != rn-1 && (i+1)%3 == 0 {
				buf.WriteByte(',')
			}
		}
	}
	return buf.String()
}
