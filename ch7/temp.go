package main

import (
	"bytes"
	"fmt"
	"io"
)

type IntSet struct { /* ... */
}

func (*IntSet) String() string {
	return "intSet"
}

func main() {
	var w io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil)
	fmt.Println(w)

	//

	/*	//var s = IntSet{}.String()  compile error: String requires *IntSet receiver
		var s IntSet
		var _ = s.String()
		var _ fmt.Stringer = &s
		//var _ fmt.Stringer = s   compile error: IntSet lacks String method
		os.Stdout.Write([] byte ( "hello" ))
		os.Stdout.Close()
		var w io.Writer
		w = os.Stdout
		w.Write([]byte("hello"))
		//w.Close() os.Stdout.Close() // OK: *os.File has Close method*/
	/*
		var any interface{}
		any = true
		any = 12.34
		any = "hello"
		any = map[string]int{"on":1}
		any = new(bytes.Buffer)
		fmt.Println(any)*/
	/*	var w io.Writer
		w = os.Stdout
		w = new(bytes.Buffer)
		w = time.Second
		println(w)
		var rwc io.ReadWriteCloser
		rwc = os.Stdout
		rwc = new (bytes.Buffer)
		println(rwc)
		w=rwc
		rwc=w*/
}
