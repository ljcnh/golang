package main

import (
	"fmt"
	"unsafe"
)

// 64-bit  32-bit
//struct{ bool; float64; int16 } // 3 words 4words
//struct{ float64; int16; bool } // 2 words 3words
//struct{ bool; int16; float64 } // 2 words 3words

func main() {
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"
}
