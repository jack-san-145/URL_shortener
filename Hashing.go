package main

import (
	"crypto/sha256"
	"fmt"
)

func hashing() {
	var hashTable [100]Node = [100]Node{}
	value := "santhosh"
	result := sha256.Sum256([]byte(value))
	fmt.Println(result)
	index := (result) % 100
	hashTable[index]
}
