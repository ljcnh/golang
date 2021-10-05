package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordsCounter int

type LinesCounter int

// 关于 return len(p)
// bytecounter的例子是这么写的 我也这么写吧...

func (c *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

func (c *LinesCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

func main() {
	var c WordsCounter
	c.Write([]byte("Hello world! 你好 世界"))
	fmt.Println(c)
	var t LinesCounter
	t.Write([]byte(`hello
qwer 
world`))
	fmt.Println(t)
}
