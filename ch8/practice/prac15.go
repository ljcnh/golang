package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	msg chan<- string
	who string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				go func(m chan<- string) {
					m <- msg
				}(cli.msg)
			}
		case cli := <-entering:
			clients[cli] = true
			cli.msg <- func() string {
				whos := make([]string, len(clients))
				for cli := range clients {
					whos = append(whos, cli.who)
				}
				return fmt.Sprintf("current clients: %v", whos)
			}()
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func handleConn(conn net.Conn) {
	fmt.Fprintf(conn, "Enter your name: ")
	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text()

	ch := make(chan string)
	go clientWriter(conn, ch)

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	inputtext := make(chan string)
	go func() {
		for input.Scan() {
			inputtext <- input.Text()
		}
	}()
	for {
		select {
		case text := <-inputtext:
			messages <- who + ": " + text
		//	这里设置的20s
		case <-time.After(20 * time.Second):
			leaving <- client{ch, who}
			messages <- who + " has left"
			conn.Close()
			return
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
