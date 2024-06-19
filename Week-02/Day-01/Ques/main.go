package main

import (
	"fmt"
)

func main() {
	list := &LinkedList{}

	// Insertion at the Beginning:
	fmt.Println("Insert Element At Beginning: ")
	list.InsertionAtBeginning(7)
	list.InsertionAtBeginning(21)
	list.InsertionAtBeginning(69)
	list.InsertionAtBeginning(22)
	list.traverse()

	// Insertion at the End:
	fmt.Println("Insert Element At the End:")
	list.InsertionAtEnd(44)
	list.InsertionAtEnd(89)

	list.traverse()

	// Insertion at a Specific Position:
	position := 2
	data := 13
	fmt.Printf("Insert data at position: %d\n", data)
	list.InsertAtPosition(position, data)

	list.traverse()

	// Deletion of a Node:
	fmt.Println("Deletion of the Node: ")
	list.deleteNode(13)
	list.traverse()

	// Deletion at the Beginning:
	fmt.Println("Deletetion from the Beginning: ")
	list.DeleteFirst()
	list.traverse()

	// Delete at end:
	fmt.Println("Delete Node from the End: ")
	list.DeleteAtEnd()
	list.traverse()

	// Traversal:
	fmt.Println("Traversing Elements: ")
	list.Traversal()

	// Finding the Length:
	fmt.Println(list.FindLenght())

	// Search for a Node:
	fmt.Println("Elements in the list: ")
	list.Traversal()
	element := 71
	result := list.SearchNode(element)
	if !result {
		fmt.Println("Node which you searched not found in the List")
	} else {
		fmt.Printf("Searching Node Found: %d\n", element)
	}
	// Reverse a Linked List:
	fmt.Println("Reverse the List: ")
	list.ReverseList()
	list.Traversal()

	// Detecting a Loop:
	fmt.Println(list.DetectLoop())

	// Finding the Middle Element:
	middle := list.FindMiddle()
	if middle != nil {
		fmt.Println("Middle Element: ", middle.data)
	}

	// Merging Two Sorted Lists:

	// Second sorted linked list
	list2 := &LinkedList{}
	list2.InsertionAtBeginning(7)
	list2.InsertionAtBeginning(21)
	list2.InsertionAtBeginning(69)
	list2.InsertionAtBeginning(22)

	// list1:

	list1 := &LinkedList{}
	list1.InsertionAtBeginning(7)
	list1.InsertionAtBeginning(21)
	list1.InsertionAtBeginning(3)
	list1.InsertionAtBeginning(31)
	list1.InsertionAtBeginning(2)
	list1.InsertionAtBeginning(9)

	mergeHead := MergeSort(list1.head)
	mergeList := &LinkedList{head: mergeHead}
	fmt.Print("Merged List: ")
	mergeList.PrintList()
	fmt.Println()

	// Removing Duplicates:

	fmt.Println("Original List: ")
	list.PrintList()

	list.RemoveDuplicates()
	fmt.Println("List after removing duplicates: ")

	list.PrintList()

	// Adding Two Numbers Represented by Lists:
	// First number: 342 (represented as 2 -> 4 -> 3)
	list1 = &LinkedList{}
	list1.InsertionAtEnd(2)
	list1.InsertionAtEnd(4)
	list1.InsertionAtEnd(3)

	// Second number: 465 (represented as 5 -> 6 -> 4)
	list2 = &LinkedList{}
	list2.InsertionAtEnd(5)
	list2.InsertionAtEnd(6)
	list2.InsertionAtEnd(4)

	// Adding two lists
	resultList := AddTwoLists(list1, list2)

	// Print the result list
	fmt.Print("Sum: ")
	resultList.PrintList() // Output: Sum: 7 0 8
}

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

// Insertion at the Beginning: Write a function to insert a node at the beginning of a singly linked list.
func (list *LinkedList) InsertionAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) traverse() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
		if current == nil {
			fmt.Println("nil")
			return
		}
	}
	fmt.Println()
}

// Insertion at the End: Implement a function to insert a node at the end of a singly linked list.
func (list *LinkedList) InsertionAtEnd(data int) {
	newNode := &Node{data: data}

	// If the list is empty, make the new node the head
	if list.head == nil {
		list.head = newNode
		return
	}

	// Traverse the last to find the last node
	current := list.head
	for current.next != nil {
		current = current.next
	}

	// Set the next pointer of the last node to the new node
	current.next = newNode
}

// Insertion at a Specific Position: Write a function to insert a node at a specific position in a singly linked list.
func (list *LinkedList) InsertAtPosition(position, data int) {
	// create a new node with given data
	newNode := &Node{data: data}

	// If the position is 0 or less, insert at the beginning
	if position <= 0 {
		newNode.next = list.head
		list.head = newNode
		return
	}

	// Traverse the list to find the position just before the desired position
	current := list.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}

	// If the position is out of bound, prints an error message
	if current == nil {
		fmt.Println("Position out of bounds.")
		return
	}

	// Insert new node at the desired position
	newNode.next = current.next
	current.next = newNode
}

// Deletion of a Node:
func (list *LinkedList) deleteNode(data int) {
	// If list is empty, return
	if list.head == nil {
		return
	}

	// If the node to delete is the head node
	if list.head.data == data {
		list.head = list.head.next
		return
	}

	// Traverse the list to find the node to delete and the node just before it
	prev := list.head
	current := list.head.next
	for current != nil && current.data != data {
		prev = current
		current = current.next
	}

	// If the node to delete it found, remove it
	if current != nil {
		prev.next = current.next
	}
}

// Deletion at the Beginning: Write a function to delete the first node from a singly linked list.
func (list *LinkedList) DeleteFirst() {
	if list.head == nil {
		return
	}
	list.head = list.head.next
}

// Deletion at the End: Implement a function to delete the last node from a singly linked list.
func (list *LinkedList) DeleteAtEnd() {
	// If list is empty nothing to delete
	if list.head == nil {
		return
	}

	// If list contains only one node, remove it by setting the head to nil
	if list.head.next == nil {
		list.head = nil
		return
	}

	// Traverse the list to find the second-last node
	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	// Remove the last node by setting the next pointer of the second-last node to nill
	current.next = nil
}

// Traversal: Write a function to traverse and print all elements of a singly linked list.
func (list *LinkedList) Traversal() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// Finding the Length: Implement a function to find the length of a singly linked list.
func (list *LinkedList) FindLenght() int {
	count := 0
	current := list.head
	for current != nil {
		count++
		current = current.next
	}

	return count
}

// Search for a Node: Write a function to search for a specific element in a singly linked list.
func (list *LinkedList) SearchNode(element int) bool {
	current := list.head
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

// Reverse a Linked List: Implement a function to reverse a singly linked list.
func (list *LinkedList) ReverseList() {
	var prev *Node
	current := list.head
	var next *Node

	for current != nil {
		next = current.next // store the next node
		current.next = prev // Reverse the current node's pointer
		prev = current
		current = next
	}
	list.head = prev
}

// Detecting a Loop: Write a function to detect if a singly linked list contains a loop.
func (list *LinkedList) DetectLoop() bool {
	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// Finding the Middle Element: Implement a function to find the middle element of a singly linked list.
/*
To find the middle element of a singly linked list, we can use the "Tortoise and Hare" approach,
where two pointers are used: one moves twice as fast as the other. When the faster pointer reaches the end of the list,
the slower pointer will be at the middle element.
*/
func (list *LinkedList) FindMiddle() *Node {
	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// Merging Two Sorted Lists: Write a function to merge two sorted singly linked lists into a single sorted list.
// Sort the list:
func MergeSort(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	// Split the list into two halves
	left, right := SplitList(head)

	// Recursively sort each half
	left = MergeSort(left)
	right = MergeSort(right)

	// Merge the sorted halves
	return Merge(left, right)
}

// Split list
func SplitList(head *Node) (*Node, *Node) {
	var prev *Node
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	// Break the list into two halves
	if prev != nil {
		prev.next = nil
	}
	return head, slow
}

// Merge the list
func Merge(left, right *Node) *Node {
	dummy := &Node{}
	current := dummy

	for left != nil && right != nil {
		if left.data < right.data {
			current.next = left
			left = left.next
		} else {
			current.next = right
			right = right.next
		}
		current = current.next
	}

	if left != nil {
		current.next = left
	} else {
		current.next = right
	}

	return dummy.next
}

// PrintList traverses the linked list and prints its elements
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// Removing Duplicates: Implement a function to remove duplicates from an unsorted singly linked list.
func (list *LinkedList) RemoveDuplicates() {
	if list.head == nil {
		return
	}

	seen := make(map[int]bool)
	current := list.head
	seen[current.data] = true

	for current.next != nil {
		if seen[current.next.data] {
			current.next = current.next.next
		} else {
			seen[current.next.data] = true
			current = current.next
		}
	}
}

// Adding Two Numbers Represented by Lists: Write a function to add two numbers represented by two singly linked lists.

func AddTwoLists(list1, list2 *LinkedList) *LinkedList {
	list := &LinkedList{}
	p1 := list1.head
	p2 := list2.head
	var carry int

	for p1 != nil || p2 != nil {
		var sum int
		if p1 != nil {
			sum += p1.data
			p1 = p1.next
		}
		if p2 != nil {
			sum += p2.data
			p2 = p2.next
		}
		sum += carry
		carry = sum / 10
		list.InsertionAtEnd(sum % 10)
	}

	if carry > 0 {
		list.InsertionAtEnd(carry)
	}
	return list
}
