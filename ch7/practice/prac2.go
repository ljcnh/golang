package main

import (
	"fmt"
	"io"
	"os"
)

type countWriter struct {
	writer io.Writer
	count  int64
}

func (c *countWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cW := &countWriter{writer: w}
	// io.Writer 本来就要求返回一个指针类型
	fmt.Println(&(cW.count)) // 0xc000004088 count的地址   #1
	return cW, &(cW.count)
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Println(writer)                 // &{0xc000006018 0}
	fmt.Fprint(writer, "Hello world\n") //Hello world  在这一步中，会调用countWriter.Write 所以count会改变
	fmt.Println(writer)                 // &{0xc000006018 12}
	fmt.Println(*count)

	fmt.Println(count) // 0xc000004088 count的地址  与 #1 的输出是一样
}
