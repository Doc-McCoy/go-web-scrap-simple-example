package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)
// O goquery foi instalado previamente com o comando:
// go get github.com/PuerkitoBio/goquery


func main() {

	url := "https://golangcode.com"
	blogTitles, err := GetLatestBlogTitles(url)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Blog titles:")
	fmt.Printf(blogTitles)
}


func GetLatestBlogTitles(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	titles := ""
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "- " + s.Text() + "\n"
	})

	return titles, nil
}
