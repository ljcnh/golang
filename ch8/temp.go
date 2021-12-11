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
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	msg chan<- string // an outgoing message channel
	who string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.msg <- msg
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
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	reset := make(chan struct{})
	end := make(chan struct{})
	input := bufio.NewScanner(conn)
	go func() {
		for input.Scan() {
			reset <- struct{}{}
			messages <- who + ": " + input.Text()
		}
		end <- struct{}{}
	}()
	// NOTE: ignoring potential errors from input.Err()
loop:
	for {
		select {
		case <-time.After(20 * time.Second):
			// set only 1 minute to test
			break loop
		case <-end:
			break loop
		case <-reset:
			// do nothing
		}
	}
	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//package main
//
//import (
//	"fmt"
//	"golang.org/x/net/html"
//	"log"
//	"net/http"
//	"os"
//)
//
//var tokens = make(chan struct{}, 20)
//
//func crawl(url string, cancel <-chan struct{}) []string {
//	fmt.Println(url)
//	tokens <- struct{}{} // acquire a token
//	list, err := Extract(url, cancel)
//	<-tokens // release the token
//	if err != nil {
//		log.Print(err)
//	}
//	return list
//}
//
//func main() {
//	worklist := make(chan []string)
//	cancel := make(chan struct{})
//	var n int // number of pending sends to worklist
//
//	// Start with the command-line arguments.
//	n++
//	go func() { worklist <- os.Args[1:] }()
//	go func() {
//		os.Stdin.Read(make([]byte, 1)) // read a single byte
//		close(cancel)
//	}()
//
//	// Crawl the web concurrently.
//	seen := make(map[string]bool)
//
//	for ; n > 0; n-- {
//		list := <-worklist
//		for _, link := range list {
//			if !seen[link] {
//				seen[link] = true
//				n++
//				go func(link string) {
//					worklist <- crawl(link, cancel)
//				}(link)
//			}
//		}
//	}
//}
//func Extract(url string, cancel <-chan struct{}) ([]string, error) {
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//	req.Cancel = cancel
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	if resp.StatusCode != http.StatusOK {
//		resp.Body.Close()
//		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
//	}
//
//	doc, err := html.Parse(resp.Body)
//	resp.Body.Close()
//	if err != nil {
//		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
//	}
//
//	var links []string
//	visitNode := func(n *html.Node) {
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, a := range n.Attr {
//				if a.Key != "href" {
//					continue
//				}
//				link, err := resp.Request.URL.Parse(a.Val)
//				if err != nil {
//					continue // ignore bad URLs
//				}
//				links = append(links, link.String())
//			}
//		}
//	}
//	forEachNode(doc, visitNode, nil)
//	return links, nil
//}
//
//func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
//	if pre != nil {
//		pre(n)
//	}
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		forEachNode(c, pre, post)
//	}
//	if post != nil {
//		post(n)
//	}
//}
