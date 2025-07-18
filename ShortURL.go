package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var shorted_url, original_url string

type PageData struct {
	Shorturl string
}

func shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	longURL := r.FormValue("LongURL")
	shorted_url, original_url = insertToHashtable(longURL)
	ExportToHtml(w, &shorted_url)

}

func generateShortUrl(hashedUrl *int) string {
	shortLink := "shortlink/" + strconv.Itoa(*hashedUrl)
	return shortLink
}

func ExportToHtml(w http.ResponseWriter, shortlink *string) {
	data := PageData{
		Shorturl: *shortlink,
	}
	templ, err := template.ParseFiles("HTML/index.html")
	if err != nil {
		fmt.Println("Template Parsing error")
		return
	}
	templ.Execute(w, data)
}
