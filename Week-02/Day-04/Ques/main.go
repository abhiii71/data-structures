package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
}

// Implement Circular Linked List:
func main() {
	list := &CircularLinkedList{}
	// Insert at Beginning:
	list.InsertAtBeginning(71)
	list.InsertAtBeginning(21)
	list.InsertAtBeginning(45)
	list.DisplayForward()

	// Insert at Position:
	list.InsertAtEnd(69)
	list.InsertAtEnd(82)
	list.InsertAtEnd(33)
	list.DisplayForward()

	// Insert at Position:
	list.InsertAtPosition(7, 2)
	list.DisplayForward()

	// Delete from Beginning:
	list.DeleteFromBeginning()
	fmt.Println("List after deleting from beginning: ")
	list.DisplayForward()

	// Delete from End:
	list.DeleteFromEnd()
	fmt.Println("List after deleting from end: ")
	list.DisplayForward()

	// Delete from Position:
	list.DeleteFromPosition(21, 1)
	fmt.Println("After Deleting the node: ")
	list.DisplayForward()

	// Traverse Circular Linked List:
	fmt.Println("Displaying items from forward: ")
	list.DisplayForward()

	// Search Element:
	result := list.SearchElement(69)
	if result != -1 {
		fmt.Println("Element found at index: ", result)
	} else {
		fmt.Println("Element not found in the list")
	}

	// Detect Loop:
	//  created a loop by pointing the last node's next pointer to the head
	// list.tail.next = list.head

	if HasLoop(list.head) {
		fmt.Println("The circular linked list has a loop.")
	} else {
		fmt.Println("The circular linked list doesn't have a loop")
	}

	// Find Length:
	list.DisplayForward()
	lenght := list.Length()
	fmt.Printf("Length of list: %d ", lenght)

	// Split Circular Linked List:

	list.InsertAtBeginning(23)
	list.InsertAtBeginning(53)
	list.InsertAtEnd(22)
	firstHalf, secondHalf := list.SplitCircularLinkedList()
	fmt.Println("First Half: ")
	firstHalf.DisplayForward()

	fmt.Println("Second Half: ")
	secondHalf.DisplayForward()

	// Convert to Singly Linked List:
	fmt.Println("Converted the list into Singly linked list: ")
	list.ConvertToSinglyLinkedList()
	list.DisplayForward()

	// Josephus Problem:
	n := 7
	k := 3
	for i := 1; i <= n; i++ {
		list.Append(i)
	}
	list.DisplayForward()

	// Solve the jeoseph problem

	survivor := list.JosephusProblem(k)
	fmt.Printf("The persion at position %d will be the last remaining\n", survivor)

	// Check Circular List Integrity:
	if list.checkIntegrity() {
		fmt.Println("The list integrity is intact (last node points to head).")
	} else {
		fmt.Println("The list integrity is compromised (last node does not point to head).")
	}
}

// Write a function to insert a node at the beginning of the circular linked list.
func (list *CircularLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode

		list.tail = newNode
		newNode.next = list.head
		newNode.prev = list.tail
		list.head.prev = newNode
		list.tail.next = newNode
		list.head = newNode
	}
}

/*
// DisplayForward display the list from head to tail
func (list *CircularLinkedList) DisplayForward() {
	if list.head == nil {
		fmt.Println("List is empty")
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
*/

// Insert at End:
func (list *CircularLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.prev = list.tail
		newNode.next = list.head
		list.tail.next = newNode
		list.head.prev = newNode
		list.tail = newNode
	}
}

// Insert at Position:
func (list *CircularLinkedList) InsertAtPosition(data, position int) {
	if position < 0 {
		fmt.Println("Invalid position")
		return
	}

	newNode := &Node{data: data}

	if position == 0 {
		if list.head == nil {
			list.head = newNode
			list.tail = newNode
			newNode.next = newNode
			newNode.prev = newNode
		} else {
			newNode.next = list.head
			newNode.prev = list.tail
			list.head.prev = newNode
			list.head.next = newNode
			list.head = newNode
		}
		return
	}
	current := list.head
	for i := 0; i < position-1; i++ {
		current = current.next

		if current == list.head {
			fmt.Println("Position out of bounds")
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

// Write a function to delete a node from the beginning of the circular linked list.
func (list *CircularLinkedList) DeleteFromBeginning() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		return
	}
	list.head = list.head.next
	list.tail.next = list.head
	list.head.prev = list.tail
}

// Write a function to delete a node from the end of the circular linked list.
func (list *CircularLinkedList) DeleteFromEnd() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	}

	current := list.tail.prev
	current.next = list.head
	list.tail.prev = current
	list.tail = current
}

// Write a function to delete a node from a specific position in the circular linked list.
func (list *CircularLinkedList) DeleteFromPosition(data, position int) {
	if position < 0 {
		fmt.Println("Invalid Position.")
		return
	}

	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}

	if position == 0 {
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
			return
		}

		list.head = list.head.next
		list.tail.next = list.head
		list.head.prev = list.tail
		return
	}
	current := list.head
	for i := 0; i < position-1; i++ {
		current = current.next
		if current == list.head {
			fmt.Println("Position out of bound")
			return
		}
	}

	current.next = current.next.next
	current.next.prev = current

	if current.next == list.head {
		list.tail = current
	}
}

// Write a function to traverse the circular linked list and print all elements.
func (list *CircularLinkedList) DisplayForward() {
	if list.head == nil {
		fmt.Println("List is empty")
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

// Write a function to search for an element in the circular linked list.
func (list *CircularLinkedList) SearchElement(data int) int {
	if list.head == nil {
		return -1
	}

	current := list.head
	position := 0
	for {
		if current.data == data {
			return position
		}
		current = current.next
		position++
		if current == list.head {
			break
		}
	}
	return -1
}

// Verify that the circular linked list inherently forms a loop. Detect loop
/*
To verify if a circular linked list inherently forms a loop,
you can use the Floyd's cycle-finding algorithm.
This algorithm uses two pointers, slow and fast, that traverse the list at different speeds.
If there is a loop, the fast pointer will eventually catch up to the slow pointer.
*/
func HasLoop(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow := head
	fast := head.next

	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}

		slow = slow.next
		fast = fast.next.next
	}
	return true
}

// Write a function to find the length of the circular linked list.
func (list *CircularLinkedList) Length() int {
	if list.head == nil {
		return 0
	}
	current := list.head
	lenght := -1
	for {
		current = current.next
		lenght++
		if current == list.head {
			break
		}
	}
	return lenght
}

// Write a function to split the circular linked list into two halves.
func (list *CircularLinkedList) SplitCircularLinkedList() (*CircularLinkedList, *CircularLinkedList) {
	if list.head == nil || list.head.next == list.head {
		return nil, nil
	}

	slow := list.head
	fast := list.head.next

	// Find the middle node of the circular linked list
	for fast != list.head && fast.next != list.head {
		slow = slow.next
		fast = fast.next.next
	}
	// Split the circular linked list into halves
	firstHalf := &CircularLinkedList{head: list.head, tail: slow}
	secondHalf := &CircularLinkedList{head: slow.next, tail: list.tail}

	// Make the last node of the first half point to the head of the first half
	slow.next = firstHalf.head
	firstHalf.head.prev = slow

	// Make the last node of the second  half point to the head of the second half
	secondHalf.tail.next = secondHalf.head
	secondHalf.head.prev = secondHalf.tail

	return firstHalf, secondHalf
}

// Write a function to convert the circular linked list to a singly linked list.
func (list *CircularLinkedList) ConvertToSinglyLinkedList() {
	if list.head == nil {
		return
	}
	list.tail.next = nil // Break the circular reference
}

// Solve the Josephus problem using a circular linked list.
func (list *CircularLinkedList) Remove(node *Node) {
	if list.head == nil {
		return
	}
	if list.head == node && list.tail == node {
		list.head = nil
		list.tail = nil
		return
	}
	if node == list.head {
		list.head = node.next
		list.tail.next = list.head
		list.head.prev = list.tail
		return
	}
	if node == list.tail {
		list.tail = node.prev
		list.tail.next = list.head
		list.head.prev = list.tail
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (list *CircularLinkedList) Append(value int) {
	newNode := &Node{data: value}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		newNode.next = newNode // Point to itself
	} else {
		list.tail.next = newNode
		newNode.next = list.head
		list.tail = newNode
	}
}

func (list *CircularLinkedList) JosephusProblem(k int) int {
	if list.head == nil {
		return -1
	}
	current := list.head
	prev := list.head

	for prev.next != list.head {
		prev = prev.next
	}
	for current.next != current {

		// skip k-1 nodes
		for i := 0; i < k-1; i++ {
			prev = current

			current = current.next
		}
		// Removes the k-th node
		fmt.Printf("Removing %d\n", current.data)
		prev.next = current.next
		current = prev.next
	}
	list.head = current

	return current.data
}

// Check Circular List Integrity: Write a function to check the integrity of the circular linked list (i.e., last node points to the head).
/*
To check the integrity of a circular linked list, you need to ensure that the last node's next pointer points back to the head of the list. Here's a function to achieve this:

Check if the list is empty: If the list is empty, it should return true as an empty list is technically circular.
Traverse the list: Move through each node until you reach the node whose next pointer points back to the head.
Check if the last node points to the head: If it does, the integrity is confirmed.
*/
func (list *CircularLinkedList) checkIntegrity() bool {
	if list.head == nil {
		return true // AN empty list considered to be circular
	}

	current := list.head
	for current.next != nil && current.next != list.head {
		current = current.next
	}
	return current.next == list.head
}
