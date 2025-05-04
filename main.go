package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/Url_shortener/index", loadIndex).Methods("GET")
	router.HandleFunc("/URL_shortener/shortURL", shortUrlHandler).Methods("POST")
	router.HandleFunc("/URL_shortener/searchURL", searchUrlHandler).Methods("POST")
	fmt.Println("Server is running")
	err := http.ListenAndServe(":8989", router)
	if err != nil {
		fmt.Println("Server Failure")
	}
}
