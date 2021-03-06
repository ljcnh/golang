package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.TextNode {
		if trimmed := strings.TrimSpace(n.Data); trimmed != "" {
			fmt.Printf("%*s%s\n", depth*2, "", trimmed)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	} else if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s ", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf("%s=\"%s\" ", a.Key, a.Val)
		}
		if n.FirstChild != nil {
			fmt.Printf(">\n")
		} else {
			fmt.Printf("/>\n")

		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
