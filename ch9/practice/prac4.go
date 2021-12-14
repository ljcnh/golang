package main

import (
	"fmt"
)

func f() {
	ch := make(chan struct{})
	var count int64 = 0
	fmt.Println("Number of goroutines:")
	for {
		count++
		fmt.Printf("\r%d", count)
		go func() { <-ch }()
	}
}

// 不要随便运行这个 会卡死。。。。
func main() {
	f()
}
