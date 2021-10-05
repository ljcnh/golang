package main

import (
	"fmt"
	"io"
	"strings"
)

// 按照io.LimitedReader 写的

type MyLimitReader struct {
	R io.Reader
	N int64
}

func (l *MyLimitReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r, n}
}

func main() {
	r := LimitReader(strings.NewReader("<html><body><h1>hello</h1></body></html>aaaaa"), 40)
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	buffer = buffer[:n]
	fmt.Println(n, err, string(buffer))
}
