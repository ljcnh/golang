package main

import "fmt"

func main() {
	/*	var x, y []int
			for i := 0; i < 10; i++ {
				y = appendInt(x, i)
				fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
				x = y
			}
		x := []int{0, 1, 2, 3, 4}
		y := []int{5, 6, 7, 8, 9}
		fmt.Printf("%d cap=%d\t%v\n", len(x), cap(x), x)
		x = appendInt(x, y...)
		fmt.Printf("%d cap=%d\t%v\n", len(x), cap(x), x)
	*/
	var stack []int
	v := []int{0, 1, 2, 3, 4}
	stack = append(stack, v...)
	top := stack[len(stack)-1]
	fmt.Println(top)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	fmt.Println(remove(stack, 3))
}

func appendIntF(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
