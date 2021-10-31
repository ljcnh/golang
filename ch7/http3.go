package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// Handle 第二个参数 是Handler类型
	// http.HandlerFunc 将 db.list 变为  http.HandlerFunc类型
	// 但是HandlerFunc有这么一个函数func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
	// 而是Handler类型的要求就是用一个ServeHTTP函数
	// 所以此时db.list成为 Handler类型

	// 注意db.list 是个函数
	//mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/list", db.list)
	mux.Handle("/price", http.HandlerFunc(db.price))

	// mux有一个这样的函数func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
	log.Fatal(http.ListenAndServe("localhost:12345", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
