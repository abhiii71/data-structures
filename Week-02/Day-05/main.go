package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type CircularDoublyLinkedList struct {
	head *Node
	tail *Node
}

func main() {
	fmt.Println("Circular Doubly Linked Lists")

	list := &CircularDoublyLinkedList{}

	// InsertAtBeginning:
	list.InsertAtBeginning(71)
	list.InsertAtBeginning(22)

	fmt.Print("Printing List Forward: ")
	list.PrintListForward()

	// InsertAtEnd:
	list.InsertAtEnd(13)
	list.InsertAtEnd(81)

	fmt.Print("Printing List Backward: ")
	list.PrintListBackward()

	// DeleteNode:
	list.DeleteNode(81)
	fmt.Print("List after Deleting Node: ")
	list.PrintListForward()
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (list *CircularDoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		newNode.next = newNode
		newNode.prev = newNode
		list.head = newNode
	} else {
		tail := list.head.prev
		newNode.next = list.head
		newNode.prev = tail
		list.head.prev = newNode
		tail.next = newNode
		list.head = newNode

	}
}

// InsertAtEnd inserts a new node at the end of the list
func (list *CircularDoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		newNode.next = newNode
		newNode.prev = newNode
		list.head = newNode
		list.tail = newNode
	} else {
		tail := list.head.prev
		newNode.next = list.head // newNode's next should point to head
		newNode.prev = tail      // newNode's prev points to current tail
		tail.next = newNode
		list.head.prev = newNode // Update the tail to newNode
	}
}

// DeleteNode deletes a node with specific data from the list
func (list *CircularDoublyLinkedList) DeleteNode(data int) {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	temp := list.head
	for temp.data != data {
		temp = temp.next
		if temp == list.head {
			fmt.Println("Node not found.")
			return
		}
	}

	if temp.next == temp { // list contains only one element

		list.head = nil
	} else {
		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		if temp == list.head {
			list.head = temp.next
		}
	}
}

// PrintListForward prints all the nodes in the list from head to tail
func (list *CircularDoublyLinkedList) PrintListForward() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	temp := list.head
	for {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
		if temp == list.head {
			break
		}
	}
	fmt.Println(temp.data) // Print the data of the node pointing back to the head
}

// PrintListBackward prints all the nodes in the list from tail to head
func (list *CircularDoublyLinkedList) PrintListBackward() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	temp := list.head.prev

	for {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.prev
		if temp.next == list.head.prev {
			break
		}
	}
	fmt.Println(temp.data) // Print the data of the node pointing back to the tail
}
