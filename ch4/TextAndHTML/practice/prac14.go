package main

import (
	"fmt"
	"golang/ch4/myJson"
	"html/template"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issueList").Parse(`
	<h1>{{.TotalCount}} issue</h1>
	<table>
	<tr style='text-align: left'>
	  <th>$</th>
	  <th>State</th>
	  <th>User</th>
	  <th>Title</th>
	</tr>
	{{range .Items}}
	<tr style='text-align: left'>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
 		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

func getIsssues(w http.ResponseWriter, r *http.Request) {
	allparams := r.URL.Query()
	params, ok := allparams["q"]
	if !ok {
		fmt.Fprintf(w, "param \"q\" not exist")
		return
	}
	fmt.Println("1URL", r.URL)
	fmt.Println("2URL.Path", r.URL.Path[1:])
	result, err := myJson.SearchIssues(params)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	if err := issueList.Execute(w, result); err != nil {
		fmt.Fprintf(w, "%v", err)
	}
}

//	http://localhost:8910/?q=golang html
func main() {
	http.HandleFunc("/", getIsssues)
	log.Fatal(http.ListenAndServe("localhost:8910", nil))
}
