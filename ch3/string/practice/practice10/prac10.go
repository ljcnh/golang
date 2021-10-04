package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "23465876qw36324625fdgdgfsd"
	fmt.Println(comma(s))
	//a := []byte(s)
	fmt.Println(len(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if i != n-1 && (n-i-1)%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
