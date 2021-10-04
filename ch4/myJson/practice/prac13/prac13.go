package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const titleURL = "http://www.omdbapi.com/?apikey=2d355102&t="

type Movie struct {
	Response string
	Error    string
	Title    string
	Year     string
	Writer   string
	Actors   string
	Poster   string
}

func main() {
	q := strings.Join(os.Args[1:], " ")
	if q == "" {
		fmt.Printf("Error: 没有查询输入")
		os.Exit(1)
	}
	resp, err := http.Get(titleURL + q)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	if movie.Response == "False" {
		log.Fatal(movie.Error)
	}
	fmt.Printf("Found movie %#v (%s)\n", movie.Title, movie.Year)
	if movie.Poster == "N/A" {
		fmt.Print("Error: 该电影没有海报")
		os.Exit(1)
	}
	poster, err := http.Get(movie.Poster)
	if err != nil {
		log.Fatal(err)
	}
	defer poster.Body.Close()
	img := bufio.NewReader(poster.Body)
	filename := movie.Title + ".png"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	_, err = io.Copy(writer, img)
	if err != nil {
		file.Close()
		os.Remove(filename)
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("结束")
}
