package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq()
}

func wordfreq() {
	var counts = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		line := in.Text()
		counts[line]++
	}
	for i, v := range counts {
		fmt.Printf("%s\t%d\n", i, v)
	}
}
