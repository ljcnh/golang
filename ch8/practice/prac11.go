package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		return
	}
	log.Printf("fastest URL is %s", fetch(os.Args[1:]))
}

func fetch(urls []string) string {
	response := make(chan string)
	done := make(chan struct{})
	defer close(done)
	defer close(response)
	for _, url := range urls {
		go func(url string) {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			req.Cancel = done
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				return
			}
			done <- struct{}{}
			response <- url
		}(url)
	}
	return <-response
}
