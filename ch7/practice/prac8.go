package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}
func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}
func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}
func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}
func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

var fields = [...]string{"None", "Title", "Artist", "Album", "Year", "Length"}

func main() {
	var clickFirst, clickSecond int
	var curClick int
	var isAsc [6]bool
	for true {
		fmt.Println("\n 'Title': click 1, 'Artist': click 2,")
		fmt.Printf(" 'Album': click 3, 'Year': click 4, 'Length':click 5, others to exit\n")
		fmt.Scanln(&curClick)
		if curClick < 0 || curClick > 5 {
			os.Exit(1)
		}
		isAsc[curClick] = !isAsc[curClick]
		if clickFirst != curClick {
			if clickFirst != 0 {
				clickSecond = clickFirst
			}
			clickFirst = curClick
		}
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			switch clickFirst {
			case 1:
				if x.Title != y.Title {
					return (x.Title < y.Title) == isAsc[1]
				}
			case 2:
				if x.Artist != y.Artist {
					return (x.Artist < y.Artist) == isAsc[2]
				}
			case 3:
				if x.Album != y.Album {
					return (x.Album < y.Album) == isAsc[3]
				}
			case 4:
				if x.Year != y.Year {
					return (x.Year < y.Year) == isAsc[4]
				}
			case 5:
				if x.Length != y.Length {
					return (x.Length < y.Length) == isAsc[5]
				}
			}
			switch clickSecond {
			case 1:
				if x.Title != y.Title {
					return (x.Title < y.Title) == isAsc[1]
				}
			case 2:
				if x.Artist != y.Artist {
					return (x.Artist < y.Artist) == isAsc[2]
				}
			case 3:
				if x.Album != y.Album {
					return (x.Album < y.Album) == isAsc[3]
				}
			case 4:
				if x.Year != y.Year {
					return (x.Year < y.Year) == isAsc[4]
				}
			case 5:
				if x.Length != y.Length {
					return (x.Length < y.Length) == isAsc[5]
				}
			}
			return false
		}})
		var order1, order2 []byte // show ascending or descending order
		order1, order2 = []byte("↓"), []byte("↓")
		if isAsc[clickFirst] {
			order1 = []byte("↑")
		}
		if isAsc[clickSecond] {
			order2 = []byte("↑")
		}
		fmt.Println("\nsort 1:", fields[clickFirst], string(order1), " sort 2:", fields[clickSecond], string(order2))
		printTracks(tracks)
	}

}
