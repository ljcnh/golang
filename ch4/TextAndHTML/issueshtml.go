package main

import (
	"fmt"
	"golang/ch4/myJson"
	"html/template"
	"log"
	"os"
)

var issueList = template.Must(template.New("issueshtml").Parse(`
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

func main() {
	result, err := myJson.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("report final")
}
