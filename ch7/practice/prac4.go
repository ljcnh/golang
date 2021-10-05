package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
)

// 按照 strings.Reader 写的

type MyReader struct {
	s string
	i int64
}

func (r *MyReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) *MyReader {
	return &MyReader{s, 0}
}

func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
	r := NewReader("123")
	s, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(s)
	}
	fmt.Println(string(s))
}
