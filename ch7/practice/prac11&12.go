package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.del)
	log.Fatal(http.ListenAndServe("localhost:12345", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

var mux sync.RWMutex

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<table>
			{{range $k, $v := .}}
				<tr>
					<td>{{ $k}}</td>
					<td>{{ $v}}</td>
				</tr>
			{{end}}
		</table>
	</body>
</html>`

func (db database) list(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	temp, err := template.New("list").Parse(tpl)
	if err != nil {
		panic(err)
	}
	if err := temp.Execute(w, db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%q already exists!\n", item)
		return
	}
	pricestr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(pricestr, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "price %s error", pricestr)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s create", item, pricestr)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 400
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	pricestr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(pricestr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%q nvalid\n", pricestr)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s update.", item, pricestr)
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 400
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s delete", item)
}
