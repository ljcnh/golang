package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var s2mu sync.Mutex
var s2count int

func main() {
	http.HandleFunc("/", s2handler)
	http.HandleFunc("/count", s2counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func s2handler(w http.ResponseWriter, r *http.Request) {
	s2mu.Lock()
	s2count++
	s2mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func s2counter(w http.ResponseWriter, r *http.Request) {
	s2mu.Lock()
	fmt.Fprintf(w, "Count %d\n", s2count)
	s2mu.Unlock()
}
