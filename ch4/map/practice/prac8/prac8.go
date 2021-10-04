package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var letterClass, numberClass, otherClass int
	invalid := 0
	counts := make(map[rune]bool)
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		_, ok := counts[r]
		if ok {
			continue
		}
		if unicode.IsLetter(r) {
			letterClass++
		} else if unicode.IsNumber(r) {
			numberClass++
		} else {
			otherClass++
		}
		counts[r] = true
	}
	fmt.Printf("rune\tclassCount\n")
	fmt.Printf("letter: %v,\t number,%v\t other: %v\n", letterClass, numberClass, otherClass)
	// 当然 也可以分别记录 letter，number，other 各自出现的rune
	// 1 counts := make(map[rune]int)  int: 1 letter 2:number 3:other
	// 2 再来个map 用来存储其他
	fmt.Println("所有出现过的run（包括letter，number，other）")
	for r := range counts {
		fmt.Printf("%q ", r)
	}
	fmt.Printf("\n")
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
