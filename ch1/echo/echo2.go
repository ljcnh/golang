package main

import (
	"fmt"
	"os"
)

func main() {
	for id, value := range os.Args[:] {
		fmt.Println(id, " ", value)
	}
}
