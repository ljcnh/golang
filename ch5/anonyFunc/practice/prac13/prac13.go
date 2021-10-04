package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//  已暂停 太慢
func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(link string) []string {
	fmt.Println(link)
	list, err := Extract(link)
	if err != nil {
		log.Print(err)
	}
	list = selectSameDomain(link, list)
	for _, l := range list {
		save(l)
	}
	return list
}

func selectSameDomain(link string, list []string) []string {
	srcURL, err := url.Parse(link)
	if err != nil {
		log.Print(err)
		return nil
	}
	var filted []string
	for _, l := range list {
		crawlURL, err := url.Parse(l)
		if err != nil {
			log.Print(err)
			return nil
		}
		if crawlURL.Host == srcURL.Host {
			filted = append(filted, l)
		}
	}
	return filted
}

func save(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Print(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Print(err)
		return
	}
	theURL, err := url.Parse(link)
	if err != nil {
		log.Print(err)
		return
	}
	dir := filepath.Join("./", theURL.Host, filepath.Clean(theURL.Path))
	filename := "content"
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}
	err = ioutil.WriteFile(filepath.Join(dir, filename), b, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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
