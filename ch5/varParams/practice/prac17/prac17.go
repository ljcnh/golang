package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

//	go run .\prac17.go https://allennlp.org/ img h1
func main() {
	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "usage: go run ./prac17.go https://allennlp.org/ img h1")
		os.Exit(1)
	}
	node, err := fetch(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	nodes := ElementsByTagName(node, os.Args[2:]...)
	for _, n := range nodes {
		fmt.Printf("%#v\n", n)
	}
}

func fetch(rawurl string) (*html.Node, error) {
	resp, err := http.Get(rawurl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return html.Parse(resp.Body)
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, s := range names {
			if doc.Data == s {
				nodes = append(nodes, doc)
				break
			}
		}
	}
	for it := doc.FirstChild; it != nil; it = it.NextSibling {
		t := ElementsByTagName(it, names...)
		if len(t) > 0 {
			nodes = append(nodes, t...)
		}
	}
	return nodes
}
