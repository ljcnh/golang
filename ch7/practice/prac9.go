package main

import (
	"log"
	"net/http"
	"net/url"
	"sort"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var t = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type webSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x webSort) Len() int {
	return len(x.t)
}

func (x webSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x webSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func trackToWebSort(t []*Track) webSort {
	return webSort{t, func(_, _ *Track) bool {
		return false
	}}
}

func sortByQ(t []*Track, q url.Values) {
	s := trackToWebSort(t)
	for _, key := range q["sort[]"] {
		switch key {
		case "Title":
			s = webSort{s.t, func(x, y *Track) bool { return x.Title < y.Title }}
		case "Artist":
			s = webSort{s.t, func(x, y *Track) bool { return x.Artist < y.Artist }}
		case "Album":
			s = webSort{s.t, func(x, y *Track) bool { return x.Album < y.Album }}
		case "Year":
			s = webSort{s.t, func(x, y *Track) bool { return x.Year < y.Year }}
		case "Length":
			s = webSort{s.t, func(x, y *Track) bool { return x.Length < y.Length }}
		}
	}
	sort.Sort(s)
}

//type data struct {
//	t []*Track
//	r *http.Request
//}
type data struct {
	T []*Track
	r *http.Request
}

func (d *data) NewURL(sortKey string) *url.URL {
	u := *d.r.URL
	q := u.Query()
	q.Add("sort[]", sortKey)
	u.RawQuery = q.Encode()
	return &u
}

var trackTable = template.Must(template.New("tracktable").Parse(`
<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
  <th><a href='{{.NewURL "Title"}}'>Title</a></th>
  <th><a href='{{.NewURL "Artist"}}'>Artist</a></th>
  <th><a href='{{.NewURL "Album"}}'>Album</a></th>
  <th><a href='{{.NewURL "Year"}}'>Year</a></th>
  <th><a href='{{.NewURL "Length"}}'>Length</a></th>
</tr>
{{range .T}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

// 有点小问题...
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sortByQ(t, r.URL.Query())
		trackTable.Execute(w, &data{t, r})
	})
	log.Fatal(http.ListenAndServe("localhost:12321", nil))
}
