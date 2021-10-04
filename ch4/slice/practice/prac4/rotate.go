package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(slice []int, n int) []int {
	// 注意 浅拷贝 和 深拷贝
	temp := make([]int, len(slice))
	copy(temp[0:], slice[n:])
	copy(temp[len(slice)-n:], slice[:n])
	copy(slice, temp)
	return slice
}
