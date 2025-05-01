package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var FinalURL string

func searchUrl(shortedURL *string, originalURL *string) {
	re := regexp.MustCompile(`\d+`)
	match_string := re.FindString(*shortedURL)
	match_int, _ := strconv.Atoi(match_string)
	index := match_int % 100
	fmt.Println("index", index)
	searchFromHashTable(index, originalURL)
	//redirct the Final url
	fmt.Println("Final URl - ", FinalURL)

}

func searchFromHashTable(index int, originalURL *string) {
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
		FinalURL = search(linklist, originalURL)
	}

}

func search(list *LinkedList, url *string) string {
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
