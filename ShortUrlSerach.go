package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

var FinalURL string

func searchUrlHandler(w http.ResponseWriter, r *http.Request) {
	shorted_url_map := mux.Vars(r)
	url_code, err := strconv.Atoi(shorted_url_map["url_code"])
	fmt.Println("url_code - ", url_code)
	if err != nil {
		fmt.Println("error while converting str -> int ")
		return
	}
	index := url_code % 100
	fmt.Println("index allocated - ", index)
	searchFromHashTable(index)
	fmt.Println("FinalURL - ", FinalURL)
	http.Redirect(w, r, FinalURL, http.StatusSeeOther)

}

func searchFromHashTable(index int) {
	var linklist *LinkedList
	fmt.Println(hashTable[index])
	if hashTable[index] == nil {
		FinalURL = ""
	} else if reflect.TypeOf(hashTable[index]) == reflect.TypeOf("string") {
		if v, ok := hashTable[index].(string); ok {
			FinalURL = v
		}

	} else if (reflect.TypeOf(hashTable[index]) == reflect.TypeOf(LinkedList{})) {
		if list, ok := hashTable[index].(LinkedList); ok {
			linklist = &list
		}
		FinalURL = search(linklist)
	}

}

func search(list *LinkedList) string {
	url := &original_url
	current := list.head
	list.displayLinkedList()
	for current != nil {
		if current.data == *url {
			return current.data
		} else {
			current = current.next
		}
	}
	return "Empty url"
}
