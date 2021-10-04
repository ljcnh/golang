package main

import "fmt"

func main() {
	fmt.Println(f(6))
}

func f(x int) (res int) {
	type st struct {
		y int
	}
	t := st{
		x,
	}
	defer func() {
		switch p := recover(); p {
		case t:
			res = t.y * t.y
		default:
			panic(p)
		}
	}()
	panic(t)
}
