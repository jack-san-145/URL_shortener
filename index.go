package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func loadIndex(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("HTML/index.html")
	if err != nil {
		fmt.Println("Error loading index.html")
		return
	}
	file.Execute(w, nil)
}
