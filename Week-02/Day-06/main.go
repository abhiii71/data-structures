package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type HeaderLinkedList struct {
	header *Node
}

func main() {
	fmt.Println("Header Linked Lists and Problem-Solving")

	list := NewHeaderLinkedList()

	// InsertAtEnd:
	list.InsertAtEnd(71)

	// InsertAtBeginning:
	list.InsertAtBeginning(69)
	list.InsertAtBeginning(13)

	fmt.Print("Original List: ")
	list.PrintList()

	// DeleteNode
	fmt.Println("Deleting the node '13': ")
	list.DeleteNode(13)

	// PrintList:
	list.PrintList()

	// MergeSortedLists:

	// Creating two sorted lists for merging
	list1 := NewHeaderLinkedList()
	list1.InsertAtEnd(1)
	list1.InsertAtEnd(3)
	list1.InsertAtEnd(5)

	list2 := NewHeaderLinkedList()
	list2.InsertAtEnd(2)
	list2.InsertAtEnd(4)
	list2.InsertAtEnd(6)

	mergeList := MergeSortedLists(list1, list2)
	fmt.Println("Merged sorted List: ")
	mergeList.PrintList()

	// RemoveDuplicates:
	duplicateList := NewHeaderLinkedList()
	duplicateList.InsertAtEnd(1)
	duplicateList.InsertAtEnd(1)
	duplicateList.InsertAtEnd(2)
	duplicateList.InsertAtEnd(3)
	duplicateList.InsertAtEnd(3)
	RemoveDuplicates(duplicateList)
	fmt.Println("List after removing duplicates: ")
	duplicateList.PrintList()

	// Adding two numbers represented by linked lists
	numList1 := NewHeaderLinkedList()
	numList1.InsertAtEnd(2)
	numList1.InsertAtEnd(4)
	numList1.InsertAtEnd(3)

	numList2 := NewHeaderLinkedList()
	numList2.InsertAtEnd(5)
	numList2.InsertAtEnd(6)
	numList2.InsertAtEnd(4)

	sumList := AddTwoNumbers(numList1, numList2)
	fmt.Println("Sum of two numbers represented by linked lists: ")
	sumList.PrintList()
}

// NewHeaderLinkedList initializes a new header linked list
func NewHeaderLinkedList() *HeaderLinkedList {
	return &HeaderLinkedList{header: &Node{}} // Header node with no data
}

// InsertAtEnd inserts a new node at the end of the list
func (list *HeaderLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	temp := list.header

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (list *HeaderLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data, next: list.header.next}
	list.header.next = newNode
}

// DeleteNode deletes a node with specific data from the list
func (list *HeaderLinkedList) DeleteNode(data int) {
	temp := list.header
	for temp.next != nil && temp.next.data != data {
		temp = temp.next
	}
	if temp.next == nil {
		fmt.Println("Node not found")
		return
	}
	temp.next = temp.next.next
}

// PrintList prints all the nodes in the list
func (list *HeaderLinkedList) PrintList() {
	temp := list.header.next
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")
}

// Merging two sorted linked lists
func MergeSortedLists(list1, list2 *HeaderLinkedList) *HeaderLinkedList {
	result := NewHeaderLinkedList()

	p1, p2 := list1.header.next, list2.header.next
	current := result.header

	for p1 != nil && p2 != nil {
		if p1.data <= p2.data {
			current.next = p1
			p1 = p1.next
		} else {
			current.next = p2
			p2 = p2.next
		}
		current = current.next
	}

	if p1 != nil {
		current.next = p1
	} else {
		current.next = p2
	}

	return result
}

// Removing duplicates from a sorted linked list
func RemoveDuplicates(list *HeaderLinkedList) {
	current := list.header.next
	for current != nil && current.next != nil {
		if current.data == current.next.data {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

// Adding two numbers represented by linked lists
func AddTwoNumbers(list1, list2 *HeaderLinkedList) *HeaderLinkedList {
	result := NewHeaderLinkedList()
	p1, p2 := list1.header.next, list2.header.next
	carry := 0
	current := result.header

	for p1 != nil || p2 != nil {
		sum := carry
		if p1 != nil {
			sum += p1.data
			p1 = p1.next
		}
		if p2 != nil {
			sum += p2.data
			p2 = p2.next
		}
		carry = sum / 10
		current.next = &Node{data: sum % 10}
		current = current.next
	}
	if carry > 0 {
		current.next = &Node{data: carry}
	}
	return result
}
