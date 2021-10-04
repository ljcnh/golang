package main

import (
	"fmt"
)

func tempDirs() []string {
	var r []string
	r = append(r, "a")
	r = append(r, "b")
	r = append(r, "c")
	r = append(r, "d")
	return r
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d)=%d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { fmt.Println(result + x) }()
	return double(x)
}

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//fmt.Println("triple", triple(4)) // "12"
/*	var rmdirs []func()
		for _, d := range tempDirs() {
			dir := d
			fmt.Print(d, " ", &d, " ")
			fmt.Println(dir, &dir)
			rmdirs = append(rmdirs, func() {
				fmt.Print(d, " ", &d, " ")
				fmt.Println(dir, &dir)
			})
		}
		for _, rmdir := range rmdirs {
			rmdir()
		}
		// 输出
		//a 0xc000048230 a 0xc000048250
		//b 0xc000048230 b 0xc000048280
		//c 0xc000048230 c 0xc0000482c0
		//d 0xc000048230 d 0xc0000482f0
		//d 0xc000048230 a 0xc000048250
		//d 0xc000048230 b 0xc000048280
		//d 0xc000048230 c 0xc0000482c0
		//d 0xc000048230 d 0xc0000482f0

		que := list.New()
		que.PushBack("sds")
		s := que.Front()
		fmt.Println(s.Value)
		que.Remove(s)
		fmt.Println(s.Value)
		link := "https://github.com/yyBeta/gopl/blob/master/5%E5%87%BD%E6%95%B0/ex13.go"
		theURL, _ := url.Parse(link)
		fmt.Println(theURL)
		fmt.Println(filepath.Clean(theURL.Path))
		fmt.Println(theURL.Path)
		fmt.Println(theURL.Host)
		fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
		fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
		fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
		f := square
		fmt.Println(f(3)) // "9"
		f = negative
		fmt.Println(f(3)) // "-3"

		fmt.Printf("%T\n", f) // "func(int) int"

		//f := product

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func add1(r rune) rune     { return r + 1 }
*/
