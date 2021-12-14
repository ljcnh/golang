package main

import (
	"fmt"
	"golang/ch9/memo/memo1"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 非并发
func main1() {
	m := memo1.New(httpGetBody)
	incomingURLs := []string{"a"}
	for _, url := range incomingURLs {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

// 并发 这种并发会报错  有问题
func main2() {
	m := memo1.New(httpGetBody)
	var n sync.WaitGroup
	incomingURLs := []string{"temp"}
	for _, url := range incomingURLs {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

// 并发安全  memo2
