package main

import "fmt"

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	data int
	prev *Node
	next *Node
}

func main() {
	list := &DoublyLinkedList{}

	// Insert at Beginning:
	list.InsertAtBeginning(7)
	fmt.Println("List after inserting int the beginning, display forward: ")
	list.DisplayForward()

	// Insert at End:
	list.InsertAtEnd(69)
	list.InsertAtEnd(21)
	list.InsertAtEnd(13)
	fmt.Println("List after inserting int the end, display forward: ")
	list.DisplayForward()

	// Insert at Position:
	fmt.Println("Inserting 5 at position 0:")
	list.InsertAtPosition(5, 0)
	list.DisplayForward()

	// Delete from Beginning:
	list.DeleteFromBeginning()
	fmt.Println("List after deleting from the beginning, display forward: ")
	list.DisplayForward()

	// Delete from End:
	list.DeleteFromEnd()
	fmt.Println("List after deleting from the end, display forward: ")
	list.DisplayForward()

	// Delete from Position:
	list.DeleteFromPosition(2)
	fmt.Println("List After deleting position 2, displayed forward: ")
	list.DisplayForward()

	fmt.Println("List displayed forward: ")
	list.DisplayForward()

	fmt.Println("List displayed backward: ")
	list.DisplayBackward()

	// Traverse Forward:
	fmt.Println("List Traversing Forward: ")
	list.TraverseForward()

	// Traversing Backward:
	fmt.Println("List Traversing Backward: ")
	list.TraverseBackward()

	// Search Element:
	result := list.SearchElement(7)

	if result {
		fmt.Println("Element found in the list")
	} else {
		fmt.Println("Element not found in the list")
	}

	// Reverse Doubly Linked List:
	fmt.Println("Reversing the list: ")
	list.Reverse()
	list.TraverseForward()

	// Find Middle Element:
	middle := list.FindMiddle()
	if middle != nil {
		fmt.Println("Middle Element: ", middle.data)
	} else {
		fmt.Println("List is empty")
	}

	// Remove Duplicates:
	list.InsertAtEnd(21)
	list.InsertAtEnd(13)
	list.InsertAtEnd(21) // Duplicate
	list.InsertAtEnd(13) // Duplicate
	fmt.Println("List before removing duplicates: ")
	list.DisplayForward()

	list.RemoveDuplicates()

	fmt.Println("List after removing duplicates: ")
	list.DisplayForward()

	// Sort Doubly Linked List:
	list.Sort()
	fmt.Println("List after sorting: ")
	list.DisplayForward()

	// Clone Doubly Linked List:
	clonedList := list.Clone()
	fmt.Println("Cloned list display forward: ")
	clonedList.DisplayForward()
}

// Insert at Beginning: Write a function to insert a node at the beginning of the doubly linked list.
// InsertAtBeginning inserts a new node at the beginning of the list
func (list *DoublyLinkedList) InsertAtBeginning(data int) {
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

// Implement Doubly Linked List: Create a basic doubly linked list structure in Go.
func (list *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

// Insert at Position: Write a function to insert a node at a specific position in the doubly linked list.
func (list *DoublyLinkedList) InsertAtPosition(data int, position int) {
	newNode := &Node{data: data}

	if position <= 0 {
		// If position is less than or equal to 0, insert at the beginning
		list.InsertAtBeginning(data)
		return
	}
	current := list.head
	index := 0

	// Traverse the list to find the position
	for current != nil && index < position {
		current = current.next
		index++
	}

	if current == nil {
		// if current is nil, position is beyond the end of the list, insert at the end
		list.InsertAtEnd(data)
	} else {
		// Insert the new node at the found position
		newNode.next = current
		newNode.prev = current.prev
		if current.prev != nil {
			current.prev.next = newNode
		} else {
			// If inserting at the head position
			list.head = newNode
		}
		current.prev = newNode
	}
}

// Delete from Beginning: Write a function to delete a node from the beginning of the doubly linked list.
/*
To delete a node from the beginning of a doubly linked list,
you need to update the head pointer to point to the next node and ensure that the previous pointer of the new head is set to `nil`.
If the list becomes empty after the deletion, you should also update the tail pointer to `nil`.
*/
func (list *DoublyLinkedList) DeleteFromBeginning() {
	if list.head == nil {
		fmt.Println("List is empty,nothing to delete")
		return
	}

	if list.head == list.tail {
		// If there's only one node in the list
		list.head = nil
		list.tail = nil
	} else {
		// Move the head to the next node
		list.head = list.head.next
		list.head.prev = nil
	}
}

// Delete from End: Write a function to delete a node from the end of the doubly linked list.
/*
To delete a node from the end of a doubly linked list,
you need to update the `tail` pointer of the list and adjust the `prev` pointer of the new tail node.
*/
func (list *DoublyLinkedList) DeleteFromEnd() {
	if list.tail == nil {
		fmt.Println("List is empty.")
		return
	}
	if list.tail.prev == nil {
		// Only 1 element in the list
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}
}

// Delete from Position: Write a function to delete a node from a specific position in the doubly linked list.
/*
To delete a node from a specific position in a doubly linked list,
you need to traverse the list to the desired position and then adjust the pointers of the adjacent nodes accordingly.
*/
func (list *DoublyLinkedList) DeleteFromPosition(position int) {
	if position < 0 {
		return // invalid position
	}

	current := list.head
	for i := 0; current != nil && i < position; i++ {
		current = current.next
	}

	if current == nil {
		return // Position is out of bounds
	}

	if current.prev != nil {
		current.prev.next = current.next
	} else {
		list.head = current.next
	}

	if current.next != nil {
		current.next.prev = current.prev
	} else {
		list.tail = current.prev
	}
}

// Traverse Forward: Write a function to traverse the doubly linked list from the head to the tail.
func (list *DoublyLinkedList) TraverseForward() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

// Traverse Backward: Write a function to traverse the doubly linked list from the tail to the head.
func (list *DoublyLinkedList) TraverseBackward() {
	current := list.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println()
}

// Search Element: Write a function to search for an element in the doubly linked list.
func (list *DoublyLinkedList) SearchElement(element int) bool {
	current := list.head
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

// Reverse Doubly Linked List: Write a function to reverse the doubly linked list.
/*To reverse a doubly linked list, you need to swap the `next` and `prev` pointers of each node.
Additionally, you need to update the `head` and `tail` pointers of the list.
*/
func (list *DoublyLinkedList) Reverse() {
	current := list.head
	var temp *Node

	// Swap next and prev for all nodes of the list
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	// Before changing head, check for the cases likes empty list and list with only one node
	if temp != nil {
		list.head = temp.prev
	}
}

// Find Middle Element: Write a function to find the middle element of the doubly linked list.
/*
To find the middle element of a doubly linked list, you can use the "tortoise and hare" approach.
This method involves using two pointers: one moves one step at a time (slow pointer),
and the other moves two steps at a time (fast pointer). When the fast pointer reaches the end of the list,
the slow pointer will be at the middle.
*/
func (list *DoublyLinkedList) FindMiddle() *Node {
	if list.head == nil {
		return nil // list is empty
	}
	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// Remove Duplicates: Write a function to remove duplicate elements from the doubly linked list.
func (list *DoublyLinkedList) RemoveDuplicates() {
	if list.head == nil {
		return
	}

	seen := make(map[int]bool)
	current := list.head

	for current != nil {
		if seen[current.data] {
			// Duplicate found, remove it
			if current.prev != nil {
				current.prev.next = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			}
			if current == list.head {
				list.head = current.next
			}
			if current == list.tail {
				list.tail = current.prev
			}
		} else {
			seen[current.data] = true
		}
		current = current.next
	}
}

// DeleteNode deletes a node from the list
func (list *DoublyLinkedList) DeleteNode(data int) {
	current := list.head
	for current != nil {
		if current.data == data {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				list.head = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				list.tail = current.prev
			}
			return
		}
		current = current.next
	}
}

// Sort Doubly Linked List: Implement a function to sort the doubly linked list.
func (list *DoublyLinkedList) Sort() {
	if list.head == nil {
		return
	}

	current := list.head.next
	for current != nil {
		key := current.data
		prev := current.prev

		// Move elements of list that are greater than key to one position ahead of their current position
		for prev != nil && prev.data > key {
			prev.next.data = prev.data
			prev = prev.prev
		}

		// Place the key at after just smaller than it.
		if prev == nil {
			list.head.data = key
		} else {
			prev.next.data = key
		}

		current = current.next
	}
}

// Clone Doubly Linked List: Write a function to create a deep copy of the doubly linked list.
func (list *DoublyLinkedList) Clone() *DoublyLinkedList {
	if list.head == nil {
		return &DoublyLinkedList{}
	}
	clonedList := &DoublyLinkedList{}
	current := list.head

	for current != nil {
		clonedList.InsertAtEnd(current.data)
		current = current.next
	}
	return clonedList
}

// DisplayForward displays the list from head to tail
func (list *DoublyLinkedList) DisplayForward() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// DisplayBackward displays the list from tail to head
func (list *DoublyLinkedList) DisplayBackward() {
	current := list.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}
