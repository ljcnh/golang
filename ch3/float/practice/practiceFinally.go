// 最终版
// http://localhost:8902/?height=600&width=1200
// http://localhost:8902/eggbox?height=600&width=1200
// http://localhost:8902/saddle?height=300&width=500

package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

var height, width float64 = 300, 600
var cells float64 = 100
var xyrange float64 = 30.0
var xyscale float64 = width / 2 / xyrange
var zscale float64 = height * 0.4
var angle float64 = math.Pi / 6
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type zFunc func(x, y float64) float64

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/eggbox", eggboxs)
	http.HandleFunc("/saddle", saddles)
	http.ListenAndServe("localhost:8902", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	updateHW(r)
	surface(w, "f")
}

func eggboxs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	updateHW(r)
	surface(w, "eggbox")
}

func saddles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	updateHW(r)
	surface(w, "saddle")
}

func updateHW(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}
	for k, v := range r.Form {
		if k == "height" {
			h, _ := strconv.ParseFloat(v[0], 64)
			if h > 0 {
				height = h
			}
		}
		if k == "width" {
			w, _ := strconv.ParseFloat(v[0], 64)
			if w > 0 {
				width = w
			}
		}
	}
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
}

func surface(w io.Writer, fname string) {
	var fn zFunc
	switch fname {
	case "saddle":
		fn = saddle
	case "eggbox":
		fn = eggbox
	default:
		fn = f
	}
	zmin, zmax := getMinMAX()
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #ff0000; fill: #0000ff; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay := corner(i+1, j, fn)
			bx, by := corner(i, j, fn)
			cx, cy := corner(i, j+1, fn)
			dx, dy := corner(i+1, j+1, fn)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				//fmt.Errorf("corner() 产生非数值")
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke:%s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i int, j int, fn zFunc) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := fn(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := 0.9 * math.Hypot(x, y)
	return math.Sin(r) / r
}

// 新增形状  3.2  eggbox和saddle
func eggbox(x, y float64) float64 {
	r := 0.1 * (math.Cos(x) + math.Cos(y))
	return r
}
func saddle(x, y float64) float64 {
	a := 25.0 * 25.0
	b := 17.0 * 17.0
	r := y*y/a - x*x/b
	return r
}

func color(i, j int, zmin, zmax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}
	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}

func getMinMAX() (min, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return min, max
}
