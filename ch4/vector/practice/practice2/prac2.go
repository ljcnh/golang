package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

var hashMethod = flag.String("s", "sha256", "default SHA256  1.SHA384  2.SHA512")

// go run .\prac3.go -s sha512
func main() {
	var str string
	fmt.Printf("您选择的是：%s\n", *hashMethod)
	fmt.Scanln(&str)
	printHash(strings.ToUpper(*hashMethod), str)
}

func printHash(flag string, str string) {
	if flag == "SHA384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
	} else if flag == "SHA512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	}
}
