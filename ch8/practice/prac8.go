package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	reset := make(chan struct{})
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for input.Scan() {
			reset <- struct{}{}
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				echo(c, s, 1*time.Second)
			}(input.Text())
		}
	}()
	for {
		select {
		case <-time.After(10 * time.Second):
			wg.Wait()
			if tc, ok := c.(*net.TCPConn); ok {
				tc.CloseWrite()
			} else {
				c.Close()
			}
			fmt.Println("10s no input, connection closed")
			return
		case <-reset:

		}
	}
	ticker.Stop()
}
