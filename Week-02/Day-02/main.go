package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func main() {
	list := LinkedList{}
	list.InsertAtEnd(71)
	list.InsertAtEnd(69)
	list.InsertAtEnd(22)
	list.InsertAtEnd(7)
	fmt.Println("Original list: ")
	list.PrintList()

	list.Reverse()
	fmt.Println("Reverse list: ")
	list.PrintList()

	// Creating a loop for testing
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)
	//	list.head.next.next.next.next = list.head.next
	list.PrintList()

	if list.DetectLoop() {
		fmt.Println("Loop Detected")
	} else {
		fmt.Println("No loop Detected")
	}

	// Finding the Middle Element of a Linked List
	middle := list.FindMiddle()
	if middle != nil {
		fmt.Printf("Middle element: %d\n", middle.data)
	} else {
		fmt.Println("The list is empty")
	}
}

// Reverse: reverse the linked list
func (list *LinkedList) Reverse() {
	var prev, next *Node
	current := list.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev

	//
	middle := list.FindMiddle()
	if middle != nil {
		fmt.Printf("Middle Element: %d\n", middle.data)
	} else {
		fmt.Println("The list is empty")
	}
}

// InsertAtEnd: inserts a new node at the end of the list
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}

	last := list.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

// PrintList  prints all the nodes in the list
func (list *LinkedList) PrintList() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")
}

// DetectLoop detects if there is a loop in the linked list
func (list *LinkedList) DetectLoop() bool {
	slow, fast := list.head, list.head
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// FindMiddle: Finds the middle element of the linked list
func (list *LinkedList) FindMiddle() *Node {
	slow, fast := list.head, list.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
