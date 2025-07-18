package main

import (
	"net/http"
	"reflect"

	"regexp"
	"strconv"

)

var FinalURL string

func searchUrlHandler(w http.ResponseWriter, r *http.Request) {
	searchUrl(&shorted_url)
	http.Redirect(w, r, FinalURL, http.StatusSeeOther)

}

func searchUrl(shortedURL *string) {
	re := regexp.MustCompile(`\d+`)
	match_string := re.FindString(*shortedURL)
	match_int, _ := strconv.Atoi(match_string)
	index := match_int % 100
	searchFromHashTable(index)

}

func searchFromHashTable(index int) {
	var linklist *LinkedList
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
