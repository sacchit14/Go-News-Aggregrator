package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Go is great</h1>")
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	defer resp.Body.Close()
	byte, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(byte, &n)
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex

	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	defer resp.Body.Close()
	byte, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(byte, &s)
	queue := make(chan News, 30)

	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)

	for element := range queue {
		for idx, _ := range element.Keywords {
			newsMap[element.Titles[idx]] = NewsMap{element.Keywords[idx], element.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
