package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "prac2:%v\n", err)
		os.Exit(1)
	}
	mp := make(map[string]int)
	mp = dfs(mp, doc)
	for node, cnt := range mp {
		fmt.Println(node, " : ", cnt)
	}
}

func dfs(mp map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return mp
	}
	if n.Type == html.ElementNode {
		mp[n.Data]++
	}
	mp = dfs(mp, n.FirstChild)
	mp = dfs(mp, n.NextSibling)
	return mp
}
