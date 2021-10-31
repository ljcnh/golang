package main

import (
	"fmt"
	"io"
	"os"
)

/*type IntSet struct {
}

func (*IntSet) String() string {
	return "intSet"
}*/
/*var unit string
var value float64

var x interface{} = []int{1, 2, 3}*/

/*const debug = true
const debug = false

func f(out io.Writer) {
	if out!=nil{
		out.Write([]byte("done\n"))
	}
}
*/

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}
func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	io.WriteString()
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
	//w = os.Stdout
	//f := w.(*os.File)
	//c := w.(*bytes.Buffer)
	//fmt.Println(f, c)
	//var w io.Writer
	//w = os.Stdout
	//rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	//w = rw                  // io.ReadWriter is assignable to io.Writer
	//w = rw.(io.Writer)      // fails only if rw == nil
	//var err error = syscall.Errno(2)
	//fmt.Println(err.Error()) // "no such file or directory"
	//fmt.Println(err)
}

/*	var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug{
		fmt.Println(buf)
	}*/

/*	var w io.Writer
	w = os.Stdout
	var v io.Writer
	v = new(bytes.Buffer)
	fmt.Println(v==w)
	fmt.Printf("%T\n", w)

	fmt.Printf("%T\n", v)*/
/*var s = "20.0"
fmt.Sscanf(s,"%f%s",&value,&unit)
fmt.Println(value)
fmt.Println(unit)*/
/*	var w io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil)
	fmt.Println(w)*/

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
