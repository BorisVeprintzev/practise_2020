package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type movie struct {
	name   string
	year   string
	contry string
}

func Parse() {
	// Request the HTML page.
	res, err := http.Get("https://www.kinopoisk.ru/film/327/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var newMovie movie
	// Find the review items
	newMovie.name = doc.Find("#headerFilm h1 span").Text()
	newMovie.year = doc.Find("#infoTable table tbody tr:nth-child(1) td:nth-child(2) div a").Text()
	newMovie.contry = doc.Find("#infoTable table tbody tr:nth-child(2) td:nth-child(2) div a").Text()
	fmt.Println(newMovie)

}

func main() {
	Parse()
}
