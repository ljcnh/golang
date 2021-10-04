package prac15

import (
	"fmt"
)

//package prac15

func main() {
	fmt.Println(min())
	fmt.Println(max(1, 45, 5, 43, 41))
	fmt.Println(max())
	fmt.Println(min(1, 45, 5, 43, 41))
}

func max(args ...int) (int, bool) {
	if len(args) == 0 {
		fmt.Printf("max:请传入参数 至少为1\n")
		return 0, false
	}
	res := args[0]
	for _, v := range args[1:] {
		if res < v {
			res = v
		}
	}
	return res, true
}

func min(args ...int) (int, bool) {
	if len(args) == 0 {
		fmt.Printf("max:请传入参数 至少为1\n")
		return 0, false
	}
	res := args[0]
	for _, v := range args[1:] {
		if res > v {
			res = v
		}
	}
	return res, true
}
