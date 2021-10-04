package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func test1(x int) int {
	res := x * 2
	fmt.Println("test2 res:", &res)
	defer func() {
		if res == 4 {
			fmt.Println("test1 res", res)
			fmt.Println("test1 res", &res)
			res = 1
		} else {
			fmt.Println("res", res)
		}
	}()
	return res
}

func test2(x int) (res int) {
	fmt.Println("test2 res:", &res)
	res = x * 2
	defer func() {
		if res == 4 {
			fmt.Println("test2 res:", res)
			fmt.Println("test2 res:", &res)
			res = 1
		} else {
			fmt.Println("res", res)
		}
	}()
	return res
}

func main() {
	fmt.Println(fetch("http://gopl.io"))
	//fmt.Println(test1(2))   4
	//fmt.Println()
	//fmt.Println(test2(2))   1
}
