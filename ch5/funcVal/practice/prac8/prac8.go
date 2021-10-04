package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var depth int
var ans *html.Node

func main() {
	ans = nil
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "usage: ./ex8.go http://example.com ID_FOR_SEARCH")
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]
	findId(url, id)
}

func findId(url string, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	node := ElementByID(doc, id)

	fmt.Printf("node: %#v\n", node)
	//fmt.Printf("Date: %s\n", node.Data)
	//fmt.Printf("Attr: %s\n", node.Attr)

	return nil
}

func ElementByID(n *html.Node, id string) *html.Node {
	forEachNode(n, id, startElement, endElement)
	return ans
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) {
	if ans != nil {
		return
	}
	if pre(n, id) {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, id, pre, post)
	}
	if post(n, id) {
		return
	}
}

func startElement(n *html.Node, id string) bool {
	if ans != nil {
		return true
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				ans = n
				fmt.Printf("depth:%d\t%s ", depth, n.Data)
				fmt.Printf("%s=\"%s\"\n", a.Key, a.Val)
				return true
			}
		}
		depth++
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		depth--
	}
	return n.FirstChild == nil
}
