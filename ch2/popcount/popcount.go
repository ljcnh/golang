package popcount

import (
	"fmt"
	"time"
)

var pc [256]byte

//var pc [256]byte = func() (pc [256]byte) {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//	return
//}()
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func TimeConsuming(tag string) func() {
	now := time.Now().UnixNano()
	return func() {
		after := time.Now().UnixNano()
		fmt.Printf("%q time cost %d ns\n", tag, after-now)
	}
}

func PopCount(x uint64) int {
	//defer TimeConsuming("PopCount")()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountCycle(x uint64) int {
	defer TimeConsuming("PopCountCycle")()
	var n int
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func PopCountByRight(x uint64) int {
	defer TimeConsuming("PopCountByRight")()
	var n int
	for i := 0; i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x >>= 1
	}
	return n
}

func PopCountByAnd(x uint64) int {
	defer TimeConsuming("PopCountByAnd")()
	var n int
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
