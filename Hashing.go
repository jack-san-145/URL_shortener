package main

import (
	"fmt"
	"reflect"
)

var hashTable [100]interface{}
var AsciiUrl int

func hash(URL string) int {

	for index := range URL {
		AsciiUrl += int(URL[index])
	}
	return AsciiUrl % 100
}

func insertToHashtable(URL string) {

	index := hash(URL)
	shortedURL := generateShortUrl(&AsciiUrl)
	fmt.Println(shortedURL)
	if hashTable[index] == nil {
		hashTable[index] = URL

	} else if reflect.TypeOf(hashTable[index]) == reflect.TypeOf("string") {
		var value string
		if v, ok := hashTable[index].(string); ok { //type assertion to get the value from the interface
			value = v
		}
		hashTable[index] = LinkedList{head: &Node{data: value}} //type assertion after making first node
		var head *LinkedList
		if list, ok := hashTable[index].(LinkedList); ok {
			head = &list
		}
		head.addNode(URL)

	} else if reflect.TypeOf(hashTable[index]) == reflect.TypeOf(LinkedList{}) {
		var head *LinkedList
		if list, ok := hashTable[index].(LinkedList); ok {
			head = &list

		} else {
			fmt.Println("Type assertion failure")
		}
		head.addNode(URL)
	}
	searchUrl(&shortedURL, &URL)

}
