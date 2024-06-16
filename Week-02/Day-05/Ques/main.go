package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type CircularDoublyLinkedList struct {
	head *Node
	tail *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func main() {
	fmt.Println("Questions over Circular Doubly Linked List: ")

	list := &CircularDoublyLinkedList{}

	// Insert at Beginning:
	list.InsertAtBeginning(71)
	fmt.Print("list after inserting at beginning: ")
	list.TraverseForward()

	// Insert at End:
	list.InsertAtEnd(69)
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)
	fmt.Print("list after inserting at end: ")
	list.TraverseForward()

	// Insert at Position:
	list.InsertAtPosition(22, 3)
	fmt.Print("list after inserting '22' at position '3': ")
	list.TraverseForward()

	// Delete from Beginning:
	list.DeletedFromBeginning()
	fmt.Print("list after deleting from beginning:")
	list.TraverseForward()

	// Delete from End:
	list.DeleteFromEnd()
	fmt.Print("list after deleting from end: ")
	list.TraverseForward()

	// Delete from Position:
	list.DeleteFromPosition(2)
	list.DeleteFromPosition(0)
	fmt.Print("list after deleting from position '0' and '2': ")
	list.TraverseForward()

	// Traverse Backward:
	fmt.Print("Traversing Backward: ")
	list.TraverseBackward()
	fmt.Println()

	// Search Element:
	result := list.SearchElement(120)

	if result {
		fmt.Println("Elemet is found in the list")
	} else {
		fmt.Println("Element is not found in the list")
	}

	// Find Length:
	length := list.FindLength()
	fmt.Println("Lenght of the list: ", length)

	// Reverse Circular Doubly Linked List:
	fmt.Print("Original list: ")
	list.TraverseForward()

	fmt.Print("Reversing list: ")
	list.ReverseCircularDoublyLinkedList()
	list.TraverseForward()
	// fmt.Print("\n")

	// Split Circular Doubly Linked List:
	list.InsertAtEnd(29)
	list.InsertAtEnd(13)
	list.InsertAtEnd(47)
	list.InsertAtEnd(33)
	list.TraverseForward()
	firstHalf, secondHalf := list.SplitCircularDoublyLinkedList()
	fmt.Print("list 1: ")
	firstHalf.TraverseForward()

	fmt.Printf("list 2: ")
	secondHalf.TraverseForward()

	// Convert to Doubly Linked List:
	list.ConvertToDoublyLinkedList()
	fmt.Println("After converting Circular -> Doubly linked list: ")
	list.TraverseForward()

	// Check Circular List Integrity:
	fmt.Println("Integrity check: ", list.CheckCircularListIntegrity())
}

// Insert at Beginning: Implement a function to insert a node at the beginning of the circular doubly linked list.
func (list *CircularDoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		newNode.next = newNode
		newNode.prev = newNode
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head // newNode's next points in the current head
		newNode.prev = list.tail // newNode's prev points in the current tail
		list.head.prev = newNode // current head's prev points to newNode
		list.tail.next = newNode // current tail's next points to newNode
		list.head = newNode      // update the head to newNode
	}
}

// Insert at End: Implement a function to insert a node at the end of the circular doubly linked list.
func (list *CircularDoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {

		newNode.next = newNode
		newNode.prev = newNode
		list.head = newNode
		list.tail = newNode
	} else {

		newNode.prev = list.tail
		newNode.next = list.head
		list.tail.next = newNode
		list.tail = newNode
	}
}

// Insert at Position: Implement a function to insert a node at a specific position in the circular doubly linked list.
func (list *CircularDoublyLinkedList) InsertAtPosition(data, position int) {
	newNode := &Node{data: data}

	if list.head == nil {
		if position != 0 {
			fmt.Println("Position out of range.")
			return
		}
		newNode.next = newNode
		newNode.prev = newNode
		list.head = newNode
		list.tail = newNode
		return
	}

	if position == 0 {
		newNode.next = list.head
		newNode.prev = list.tail
		list.head.prev = newNode
		list.tail.next = newNode
		list.head = newNode
		return
	}

	// Travese to the node just before the desired position
	current := list.head
	for i := 0; i < position-1; i++ {
		current = current.next
		if current == list.head {
			fmt.Println("Position out of range.")
			return
		}
	}
	newNode.next = current.next
	newNode.prev = current
	current.next.prev = newNode
	current.next = newNode

	if newNode.next == list.head {
		list.tail = newNode
	}
}

// Delete from Beginning: Implement a function to delete a node from the beginning of the circular doubly linked list.
func (list *CircularDoublyLinkedList) DeletedFromBeginning() {
	if list.head == nil {
		fmt.Println("List is empty")
	}

	if list.head == list.tail { // if the list has only one node
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next // Move head to the next node
		list.head.prev = list.tail // Update the previous pointer of the new head
		list.tail.next = list.head // Update the next pointer to the tail
	}
}

// Delete from End: Implement a function to delete a node from the end of the circular doubly linked list.
func (list *CircularDoublyLinkedList) DeleteFromEnd() {
	if list.head == nil {
		fmt.Println("List is empty")
	}

	if list.head == list.tail { // if the list has only one node
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev // Move tail to the prev node
		list.tail.next = list.head // Update the next pointer of the new tail
		list.head.prev = list.tail // Update the previous pointer of the head
	}
}

// Delete from Position: Implement a function to delete a node from a specific position in the circular doubly linked list.
func (list *CircularDoublyLinkedList) DeleteFromPosition(position int) {
	if list.head == nil {
		fmt.Println("List is empty")
	}

	if position == 0 { // Delete From beginning
		if list.head == list.tail {

			list.head = nil
			list.tail = nil
		} else {
			list.head = list.head.next // Move head to the next node
			list.tail.next = list.head // Update the previous pointer of the new head
			list.head.prev = list.tail // Update the next pointer of the tail
		}
		return
	}

	current := list.head
	for i := 0; i < position; i++ {
		current = current.next
		if current == list.head {
			fmt.Println("Position out of range.")

			return
		}
	}

	current.prev.next = current.next
	current.next.prev = current.prev

	if current == list.tail {
		list.tail = current.prev
	}
}

// Traverse Forward: Implement a function to traverse the circular doubly linked list in the forward direction and print all elements.
func (list *CircularDoublyLinkedList) TraverseForward() {
	if list.head == nil { // If the list is empty
		fmt.Println("List is empty.")
		return
	}
	current := list.head
	for {
		fmt.Printf("%d -> ", current.data)
		current = current.next
		if current == list.head {
			break
		}
	}
	fmt.Println()
}

// Traverse Backward: Implement a function to traverse the circular doubly linked list in the backward direction and print all elements.
func (list *CircularDoublyLinkedList) TraverseBackward() {
	if list.head == nil { // If the list is empty
		fmt.Println("List is empty.")
		return
	}

	temp := list.tail
	for {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.prev

		if temp == list.tail {
			break
		}
	}
}

// Search Element: Implement a function to search for an element in the circular doubly linked list.
func (list *CircularDoublyLinkedList) SearchElement(element int) bool {
	temp := list.head
	for {
		if temp.data == element {
			return true
		}
		temp = temp.next
		if temp == list.head {
			break
		}
	}
	return false
}

// Find Length: Implement a function to find the length of the circular doubly linked list.
func (list *CircularDoublyLinkedList) FindLength() int {
	current := list.head
	count := 0

	for {
		count++
		current = current.next
		if current == list.head {
			break
		}
	}
	return count
}

// Reverse Circular Doubly Linked List: Implement a function to reverse the circular doubly linked list.
func (list *CircularDoublyLinkedList) ReverseCircularDoublyLinkedList() {
	current := list.head
	var temp *Node
	for {
		temp = current.next
		current.next = current.prev
		current.prev = temp

		current = current.prev

		if current == list.head {
			break
		}
	}
	temp = list.head
	list.head = list.tail
	list.tail = temp
}

// Split Circular Doubly Linked List: Implement a function to split the circular doubly linked list into two halves.
func (list *CircularDoublyLinkedList) SplitCircularDoublyLinkedList() (*CircularDoublyLinkedList, *CircularDoublyLinkedList) {
	if list.head == nil || list.head.next == list.head {
		return list, nil
	}
	slow := list.head
	fast := list.head

	// Find the middle of the list
	for fast.next != list.head && fast.next.next != list.head {
		slow = slow.next
		fast = fast.next.next
	}

	// Adjust if the list has an even number of nodes
	if fast.next.next == list.head {
		fast = fast.next
	}

	// Create first half
	firstHalf := &CircularDoublyLinkedList{head: list.head, tail: slow}
	firstHalf.tail.next = firstHalf.head
	firstHalf.head.prev = firstHalf.tail

	// create second half
	secondHalf := &CircularDoublyLinkedList{head: slow.next, tail: list.tail}
	secondHalf.tail.next = secondHalf.head
	secondHalf.head.prev = secondHalf.tail

	// Break the link between the two halves
	firstHalf.tail.next = firstHalf.head
	firstHalf.head.prev = firstHalf.tail

	secondHalf.tail.next = secondHalf.head
	secondHalf.head.prev = secondHalf.tail

	// Re-establish the circular links for the two halves
	if firstHalf.tail.next == secondHalf.head {
		firstHalf.tail.next = firstHalf.head
		firstHalf.head.prev = firstHalf.tail
	}
	if secondHalf.head.prev == firstHalf.tail {
		secondHalf.tail.next = secondHalf.head
		secondHalf.head.prev = secondHalf.tail
	}

	// Disconnect the two halves
	slow.next = firstHalf.head
	slow.next.prev = secondHalf.tail

	return firstHalf, secondHalf
}

// Convert to Doubly Linked List: Write a function to convert the circular doubly linked list to a regular doubly linked list.
func (list *CircularDoublyLinkedList) ConvertToDoublyLinkedList() *DoublyLinkedList {
	if list.head == nil {
		return &DoublyLinkedList{}
	}

	dll := &DoublyLinkedList{}

	// Intiliaze the head of the doubly linked list
	current := list.head
	dll.head = &Node{data: current.data}
	dll.tail = dll.head

	// Traverse the circular list and create corresponding nodes in the doubly linked list
	current = current.next
	for current != list.head {
		newNode := &Node{data: current.data}
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
		current = current.next
	}

	// Break the circular connection
	dll.head.prev = nil

	dll.tail.next = nil
	return dll
}

// Check Circular List Integrity: Implement a function to ensure the circular doubly linked list is intact (i.e., the last node's next pointer points to the head, and the head's previous pointer points to the last node).
func (list *CircularDoublyLinkedList) CheckCircularListIntegrity() bool {
	if list.head == nil {
		return false
	}

	current := list.head
	last := list.head.prev

	// check if the  last node's next pointer points to the end
	for current != last.next {
		if current == nil {
			return false
		}

		current = current.next

	}

	// check if the  head's previous pointer points to the last node
	if last.prev != current {
		return false
	}
	return true
}
