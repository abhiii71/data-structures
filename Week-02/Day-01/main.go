package main

import "fmt"

func main() {
	// Singly linked list
	list := LinkedList{}

	list.InsertAtBeginning(3)
	list.InsertAtEnd(5)
	list.InsertAtEnd(7)
	list.InsertAtPosition(2, 1)

	fmt.Println("LinkedList after insertion: ")
	list.traverse()

	list.DeleteNode(5)
	fmt.Println("LinkedList after deletion: ")
	list.traverse()
}

// Singly Linked list:

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents linked list itself
type LinkedList struct {
	head *Node
}

// InsertAtBeginning insert a new node at the  beginning of the list
func (list *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{
		data: data,
		next: list.head,
	}
	list.head = newNode
}

// InsertAtEnd insert a new node at the end of the list
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// InsertAtPosition insert a new node at the specific position(o-based index)
func (list *LinkedList) InsertAtPosition(data, position int) {
	if position == 0 {
		list.InsertAtBeginning(data)
		return
	}

	newNode := &Node{data: data}
	current := list.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil {
		fmt.Println("Position out of bounds.")
		return
	}
	newNode.next = current.next
	current.next = newNode
}

// DeleteNode deletes the first occurrence of the given data
func (list *LinkedList) DeleteNode(data int) {
	if list.head == nil {
		return
	}
	if list.head.data == data {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	if current.next == nil {
		fmt.Println("Data not found in the list.")
		return
	}
	current.next = current.next.next
}

// Traverse prints all the elements in the LinkedList
func (list *LinkedList) traverse() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
