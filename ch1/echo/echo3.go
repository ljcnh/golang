package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	s, t := "", ""
	for _, value := range os.Args[:] {
		s += t + value
		t = " "
	}
	t2 := time.Since(t1)
	fmt.Println(s)
	fmt.Println(t2)

	t1 = time.Now()
	s = strings.Join(os.Args[:], " ")
	t2 = time.Since(t1)
	fmt.Println(s)
	fmt.Println(t2)
}
