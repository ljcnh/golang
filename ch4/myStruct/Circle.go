package myStruct

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Spokes: 23,
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 9,
		},
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)
}

/*package main

import "fmt"

type point struct {
	X, Y int
}
type circle struct {
	point
	Radius int
}
type Wheel struct {
	circle
	spokes int
}

func main() {
	w := Wheel{circle{point{8, 8}, 5}, 20}
	w = Wheel{
		spokes: 23,
		circle: circle{
			point:  point{X: 8, Y: 8},
			Radius: 9,
		},
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)
}*/
