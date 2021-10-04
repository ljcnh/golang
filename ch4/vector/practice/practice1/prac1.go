package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(countUint8(c1, c2))
	fmt.Println(countString("sadsadas", "sccadas"))
}

func countUint8(x [32]uint8, y [32]uint8) int {
	var n int
	for i := 0; i < 32; i++ {
		tx := x[i]
		ty := y[i]
		for j := 0; j < 8; j++ {
			if tx&1 != ty&1 {
				n++
			}
			tx >>= 1
			ty >>= 1
		}
	}
	return n
}

func countString(c1 string, c2 string) int {
	x := sha256.Sum256([]byte(c1))
	y := sha256.Sum256([]byte(c2))
	var n int
	for i := 0; i < len(x); i++ {
		tx := x[i]
		ty := y[i]
		for j := 0; j < 8; j++ {
			if tx&1 != ty&1 {
				n++
			}
			tx >>= 1
			ty >>= 1
		}
	}
	return n
}
