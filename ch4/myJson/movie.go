package myJson

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Movies struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
	//  开头小写
	unused int
}

var movies = []Movies{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}, unused: 2},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}, unused: 2},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}, unused: 8},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	// 输出不包含unused
	fmt.Printf("%s\n", data)
	/*
		data, err = json.MarshalIndent(movies, "", "  ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)*/

	var titles []Movies
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	// 所有的unused=0
	fmt.Println(titles)
	fmt.Println(reflect.TypeOf(titles))
}
