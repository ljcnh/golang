package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 87, 1234, 465}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(arr *[10]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
