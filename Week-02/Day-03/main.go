package main

import "fmt"

// Node represent a node in a doubly
type Node struct {
	data int
	prev *Node
	next *Node
}

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func main() {
	list := DoublyLinkedList{}
	list.InsertAtBeginnning(3)
	list.InsertAtBeginnning(71)
	list.InsertAtBeginnning(21)
	list.InsertAtEnd(69)

	//  prints all the nodes in the list from head to tail
	fmt.Println("List from head to tail: ")
	list.PrintListForward()

	//  prints all the nodes in the list from tail to head
	fmt.Println("List from tail to head: ")
	list.PrintListBackward()

	list.DeleteNode(21)
	fmt.Println("Liste after deleting node with data 21: ")
	list.PrintListForward()
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (list *DoublyLinkedList) InsertAtBeginnning(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

// InsertAtEnd inserts a new node at the end of the list
func (list *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

// DeleteNode deletes a node with specific data from the list
func (list *DoublyLinkedList) DeleteNode(data int) {
	temp := list.head
	for temp != nil && temp.data != data {
		temp = temp.next
	}
	if temp == nil {
		fmt.Println("Node not found")
		return
	}

	if temp.prev != nil {
		temp.prev.next = temp.next
	} else {
		list.head = temp.next
	}
	if temp.next != nil {
		temp.next.prev = temp.prev
	} else {
		list.tail = temp.prev
	}
}

// PrintListForward prints all the nodes in the list from head to tail
func (list *DoublyLinkedList) PrintListForward() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")
}

// PrintListBackward prints all the nodes in the list from tail to head
func (list *DoublyLinkedList) PrintListBackward() {
	temp := list.tail
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.prev
	}
	fmt.Println("nil")
}
