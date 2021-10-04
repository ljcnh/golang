package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(expand("    $abc just $test   ", f))
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$[a-zA-Z0-9]+`)
	return re.ReplaceAllStringFunc(s, f)
}

func f(s string) string {
	return s[1:] + "."
}
