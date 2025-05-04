package main

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
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = new_node
	}

}

func (list *LinkedList) displayLinkedList() {
	temp := list.head
	for temp != nil {
		temp = temp.next
	}
}
