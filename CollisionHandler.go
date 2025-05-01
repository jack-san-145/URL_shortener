package main

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) addNode(value string) {
	temp := list.head
	new_node := &Node{data: value}
	if temp == nil {
		temp = new_node
		fmt.Println(temp.data)
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = new_node
		fmt.Println(current.data)
	}

}

func (list *LinkedList) displayLinkedList() {
	temp := list.head
	for temp != nil {
		fmt.Println("->")
		fmt.Println(temp.data)
		temp = temp.next
	}
	fmt.Println("")
}
