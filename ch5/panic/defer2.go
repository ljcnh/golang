package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f2(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f2(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f2(x - 1)
}

// 函数
// defer
// Stack
// panic 报错
