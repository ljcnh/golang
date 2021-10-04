package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//点(x,y)表示复数值z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	//const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n)
		}
	}
	return color.Black
}

func getColor(n uint8) color.Color {
	paletted := [16]color.Color{
		color.RGBA{66, 30, 15, 255},    // brown 3
		color.RGBA{25, 7, 26, 255},     // violett
		color.RGBA{0, 255, 127, 255},   // 嫩绿色
		color.RGBA{4, 4, 73, 255},      // 蓝 5
		color.RGBA{0, 7, 100, 255},     // 蓝 4
		color.RGBA{12, 44, 138, 255},   // 蓝 3
		color.RGBA{24, 82, 177, 255},   // 蓝 2
		color.RGBA{57, 125, 209, 255},  // 蓝 1
		color.RGBA{134, 181, 229, 255}, // 蓝 0
		color.RGBA{211, 236, 248, 255}, // lightest 蓝
		color.RGBA{241, 233, 191, 255}, // lightest 黄
		color.RGBA{248, 201, 95, 255},  // 亮黄
		color.RGBA{240, 255, 240, 255}, // 蜜露橙
		color.RGBA{127, 255, 212, 255}, // 碧绿色
		color.RGBA{64, 224, 208, 255},  // 青绿色
		color.RGBA{240, 255, 255, 255}, // 天蓝色
	}
	return paletted[n%16]
}
