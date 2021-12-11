package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name string
	ip   string
}

func (c clock) mustCopy(dst io.Writer, src io.Reader) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Fprintf(dst, "%s: %s\n", c.name, s.Text())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "没有参数！！！")
		os.Exit(1)
	}
	clocks := make([]clock, 0)
	for _, a := range os.Args[1:] {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "没有参数 传错了！%s", a)
			os.Exit(1)
		}
		clocks = append(clocks, clock{fields[0], fields[1]})
	}
	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.ip)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go c.mustCopy(os.Stdout, conn)
	}
	for {
		time.Sleep(time.Minute)
	}
}
