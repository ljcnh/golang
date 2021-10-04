package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type comic struct {
	Num        int
	Transcript string
	Year       string
	Month      string
	Day        string
	Title      string
}

//	建立索引时 返回的Num太多 只要500个(要不然太慢了...)
const ComicNumMAX float64 = 500

func main() {
	searchTerm := strings.Join(os.Args[1:], " ")
	if _, err := os.Stat("AllInformation.json"); err != nil {
		createComicIndex()
	}

	allInfile, err := os.Open("AllInformation.json")
	if err != nil {
		log.Fatal(err)
	}
	var info comic
	if err := json.NewDecoder(allInfile).Decode(&info); err != nil {
		log.Fatal(err)
	}

	if timeOver(info) {
		createComicIndex()
	}
	var xkcd comic
	fmt.Printf("Found search term %#v in ", searchTerm)
	for i := 1; i < info.Num; i++ {
		file, err := os.Open(fmt.Sprintf("index/xkcd%d.json", i))
		if err != nil {
			if os.IsExist(err) {
				break
			}
			log.Fatal(err)
		}
		if err := json.NewDecoder(file).Decode(&xkcd); err != nil {
			file.Close()
			if i == 404 {
				continue
			}
			log.Fatal(err)
		}
		file.Close()
		if strings.Contains(strings.ToLower(xkcd.Transcript), strings.ToLower(searchTerm)) {
			fmt.Printf("https://xkcd.com/%d/\n", i)
			fmt.Println(xkcd.Transcript)
		}
	}
}

func timeOver(info comic) bool {
	nowYear := time.Now().Year()
	nowMonth := int(time.Now().Month())
	nowDay := time.Now().Day()
	year, _ := strconv.Atoi(info.Year)
	month, _ := strconv.Atoi(info.Month)
	day, _ := strconv.Atoi(info.Day)
	if nowYear > year || nowMonth > month || nowDay > day+1 {
		return true
	}
	return false
}

func createComicIndex() {
	fmt.Println("Start")
	r, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	var lastComic comic
	if err := json.NewDecoder(r.Body).Decode(&lastComic); err != nil {
		log.Fatal(err)
	}
	allInfile, err := os.OpenFile("AllInformation.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer allInfile.Close()

	allInformation, _ := json.Marshal(lastComic)
	allInfile.WriteString(string(allInformation))

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir(wd+"/index/", 0700); err != nil {
		if !os.IsExist(err) {
			log.Fatal(err)
		}
	}

	lastComic.Num = int(math.Min(float64(lastComic.Num), ComicNumMAX))

	for i := 1; i <= lastComic.Num; i++ {
		f, err := os.OpenFile(fmt.Sprintf("index/xkcd%d.json", i), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		r, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))
		if err != nil {
			f.Close()
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			r.Body.Close()
			f.Close()
			log.Fatal(err)
		}
		r.Body.Close()
		f.WriteString(string(b))
		f.Close()
	}
	fmt.Println("Complete")
}
