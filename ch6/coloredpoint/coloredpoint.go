package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	/*	red := color.RGBA{255, 0, 0, 255}
		blue := color.RGBA{0, 0, 255, 255}
		var p = ColoredPoint{&Point{1, 1}, red}
		var q = ColoredPoint{&Point{5, 4}, blue}
		fmt.Println(p.Distance(*q.Point)) // "5"
		q.Point = p.Point
		p.ScaleBy(2)
		fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"*/
	/*	p := Point{1, 2}
		q := Point{4, 6}
		distanceFromP := p.Distance        // method value
		fmt.Println(distanceFromP(q))      // "5"
		var origin Point                   // {0, 0}
		fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)
		scaleP := p.ScaleBy                // method value
		scaleP(2)                          // p becomes (2, 4)
		scaleP(3)                          // then (6, 12)
		scaleP(10)                         // then (60, 120)*/
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)
	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

type ColoredPoint struct {
	*Point
	color.RGBA
}

/*type ColoredPoint struct {
	Point
	Color color.RGBA
}*/

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.X)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.X)
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}
