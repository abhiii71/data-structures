package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func main() {
	list := LinkedList{}
	// Insertion at the Beginning:
	fmt.Println("Insert Element At Beginning: ")
	list.InsertionAtBeginning(7)
	list.InsertionAtBeginning(21)
	list.InsertionAtBeginning(69)
	list.InsertionAtBeginning(22)
	list.traverse()

	// Reverse Linked List (Iterative):
	fmt.Println("Reverse Linked List (Iterative):")
	list.ReverseListIterative()
	list.traverse()

	// Reverse Linked List (Recursive):
	fmt.Println("Reverse Linked List (Recursive):")
	list.head = list.ReverseListRecursive(list.head)
	list.traverse()

	// Detect Loop (Floyd’s Cycle-Finding Algorithm):
	// Creating a loop for testing
	// list.head.next.next.next.next = list.head.next
	if list.DetectLoop() {
		fmt.Println("Loop detected in the list.")
	} else {
		fmt.Println("No loop detected in the list.")
	}

	// Detect Loop (Hashing):
	if list.DetectLoopHashing() {
		fmt.Println("Loop detected in the list.")
	} else {
		fmt.Println("No loop detected in the list.")
	}

	// Find Middle Element (Slow and Fast Pointer):
	middle := list.FindMiddle()

	if middle != nil {
		fmt.Printf("Middle Element: %d\n", middle.data)
	} else {
		fmt.Println("The list is empty.")
	}

	// Reverse Sub-list:
	fmt.Printf("Original list:")
	list.traverse()
	start, end := 2, 4
	fmt.Printf("Reversing sub-list from position %d to %d:\n", start, end)
	list.ReverseSubList(start, end)
	fmt.Println("List after reversing sub-list: ")
	list.traverse()

	// Find Nth Node from End:
	n := 2
	node := list.FindNthFromEnd(n)
	if node != nil {
		fmt.Printf("The %dth node from the end is: %d\n", n, node.data)
	} else {
		fmt.Printf("The list has fewer than %d nodes.\n", n)
	}

	// Check Palindrome:
	if list.IsPalindrome() {
		fmt.Println("The linked list is a palindrome.")
	} else {
		fmt.Println("The linked list is not a palindrome.")
	}

	// Remove Loop:
	list.head.next.next.next.next = list.head.next
	list.DetectAndRemoveLoop()

	fmt.Println("List after removing loop: ")
	list.traverse()

	// Merge Sort on Linked List:
	list.MergeSort()
	fmt.Println("SortedList: ")
	list.traverse()

	// Partition List:
	x := 3
	list.Partition(x)
	fmt.Println("Partition List around", x, ":")
	list.traverse()

	// Rotate List: Implement a function to rotate the linked list to the right by k places.
	k := 2
	list.RotateRight(k)
	fmt.Println("Rotate List by ", k, "places: ")
	list.traverse()

	// Add Two Numbers:
	list1 := LinkedList{}
	list2 := LinkedList{}

	// First number: 342 (represented as 2 -> 4 -> 3)
	list1.InsertionAtEnd(2)
	list1.InsertionAtEnd(4)
	list1.InsertionAtEnd(3)

	// Second number: 465 (represented as 5 -> 6 -> 4)
	list2.InsertionAtEnd(5)
	list2.InsertionAtEnd(6)
	list2.InsertionAtEnd(4)

	fmt.Print("List 1: ")
	list1.PrintList()

	fmt.Print("List 2: ")
	list2.PrintList()

	result := AddTwoLists(&list1, &list2)
	// Print the result
	fmt.Print("Sum: ")
	result.PrintList()

	// Swap Nodes in Pairs:
	fmt.Print("Original List: ")
	list.PrintList()

	list.SwapPairs()

	fmt.Print("List after swapping pairs: ")
	list.PrintList()

	// Reorder List:
	fmt.Print("Original List: ")
	list.PrintList()

	list.ReorderList()

	fmt.Print("Reordered List: ")
	list.PrintList()
}

func (list *LinkedList) InsertionAtBeginning(data int) {
	newNode := &Node{data: data}

	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) traverse() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next

		if current == nil {
			fmt.Println("nil")
			return
		}

	}
	fmt.Println()
}

// Reverse Linked List (Iterative): Implement an iterative function to reverse a singly linked list.
func (list *LinkedList) ReverseListIterative() {
	var prev *Node
	current := list.head
	var next *Node

	for current != nil {
		next = current.next // store the next nodenewNode.next = list.head
		current.next = prev // Reverse the  current  node's
		prev = current
		current = next
	}
	list.head = prev
}

// Reverse Linked List (Recursive): Implement a recursive function to reverse a singly linked list.
func (list *LinkedList) ReverseListRecursive(node *Node) *Node {
	// Base case
	if node == nil || node.next == nil {
		return node
	}

	// Recuresively reverse the rest of the list
	newHead := list.ReverseListRecursive(node.next)

	// Reverse the current node's pointer
	node.next.next = node
	node.next = nil

	return newHead
}

// Detect Loop (Floyd’s Cycle-Finding Algorithm): Write a function to detect a loop in a linked list using Floyd’s Cycle-Finding Algorithm.
/*
Floyd’s Cycle-Finding Algorithm, also known as the Tortoise and Hare algorithm, is an efficient way to detect loops in a linked list.
It uses two pointers, slow and fast, that traverse the list at different speeds. If there's a loop, the fast pointer will eventually meet the slow pointer.
*/
func (list *LinkedList) DetectLoop() bool {
	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next      // move one step further
		fast = fast.next.next // move two step further
		if slow == fast {     // means there's a loop
			return true
		}
	}
	return false // If fast reaches the end, there's no loop
}

// Detect Loop (Hashing): Write a function to detect a loop in a linked list using a hashing approach.
func (list *LinkedList) DetectLoopHashing() bool {
	visited := make(map[*Node]bool)
	current := list.head

	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.next
	}
	return false
}

// Find Middle Element (Slow and Fast Pointer):
// Implement a function to find the middle element of a linked list using the slow and fast pointer approach.
/*

To find the middle element of a linked list using the slow and fast pointer approach,
you can use two pointers: slow and fast. The slow pointer moves one step at a time,
while the fast pointer moves two steps at a time. When the fast pointer reaches the end of the list,
the slow pointer will be at the middle of the list.
*/
func (list *LinkedList) FindMiddle() *Node {
	if list.head == nil {
		return nil
	}

	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

// Reverse Sub-list: Implement a function to reverse a sub-list within a given range.
func (list *LinkedList) ReverseSubList(start, end int) {
	if start >= end || list.head == nil {
		return
	}

	// Intialize pointers
	var prev, subListPrev *Node
	current := list.head

	// Traverse the list to find the start of the sub-list
	for i := 1; current != nil && i < start; i++ {
		subListPrev = current
		current = current.next
	}

	// Mark the start of the sub-list
	subListTail := current
	var subListNext *Node

	// Reverse the sub-list from start  to end
	for current != nil && start <= end {
		subListNext = current.next
		current.next = prev
		prev = current
		current = subListNext
		start++
	}

	// Reconnect the revesed sub-list with the rest of the list
	if subListPrev != nil {
		subListPrev.next = prev
	} else {
		list.head = prev
	}
	subListTail.next = current
}

// Find Nth Node from End: Write a function to find the nth node from the end of a linked list.

/*

To find the nth node from the end of a singly linked list, you can use a two-pointer technique often referred to as the "runner" or "two-pointer" approach.
This technique involves moving one pointer ahead by n nodes and then moving both pointers simultaneously until the first pointer reaches the end. At that point,
the second pointer will be at the nth node from the end.
*/

func (list *LinkedList) FindNthFromEnd(n int) *Node {
	if list.head == nil {
		return nil
	}

	first := list.head
	second := list.head

	// Move first pointer n nodes ahead
	for i := 0; i < n; i++ {
		if first == nil {
			return nil // n is greater than the number of nodes in the list
		}
		first = first.next
	}

	// Move both pointers until first reach the end
	for first != nil {
		first = first.next
		second = second.next
	}
	return second
}

// Check Palindrome: Implement a function to check if a linked list is a palindrome.
/*
A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward.
For example, "radar" and "level" are palindromic words.
*/
func (list *LinkedList) IsPalindrome() bool {
	if list.head == nil {
		return false
	}

	// Store node value in an array
	var values []int
	current := list.head
	for current != nil {
		values = append(values, current.data)
		current = current.next
	}

	// Compare values from both ends of the array
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		if values[i] != values[j] {
			return false
		}
	}
	return true
}

// Remove Loop: Write a function to remove a loop from a linked list, if present.
func (list *LinkedList) DetectAndRemoveLoop() {
	if list.head == nil || list.head.next == nil {
		return
	}

	slow := list.head
	fast := list.head

	// Detect the loop using Floyd's Cycle-Finding Algorithm
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break // Loop detected
		}
	}
	if slow != fast {
		return
	}

	// Move slow to the head and advance both slow and fast pointers until they meet at the start of the loop
	slow = list.head
	for slow.next != fast.next {
		slow = slow.next
		fast = fast.next
	}

	// Break the loop by setting the next pointer of the last node before the start of the loop to nil
	fast.next = nil
}

// Merge Sort on Linked List: Implement merge sort for a singly linked list.
func (list *LinkedList) MergeSort() {
	list.head = mergeSort(list.head)
}

/*
The MergeSort method for the LinkedList struct calls the mergeSort function with the head node.
mergeSort function recursively splits the list, sorts each half, and merges the sorted halves.
*/
func mergeSort(head *Node) *Node {
	// Base case: if the list is empty or has only one node, it's already sorted
	if head == nil || head.next == nil {
		return head
	}

	// split the list into two halves
	middle := getMiddle(head)
	nextToMiddle := middle.next
	middle.next = nil

	// Recursively sort each half
	left := mergeSort(head)
	right := mergeSort(nextToMiddle)

	// Merge the sorted halves
	sortedList := sortedMerge(left, right)
	return sortedList
}

// Finds and returns the middle node of the linked list using the slow and fast pointer approach.
func getMiddle(head *Node) *Node {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// Merges two sorted linked lists into a single sorted linked list.
func sortedMerge(left, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var result *Node
	if left.data <= right.data {
		result = left
		result.next = sortedMerge(left.next, right)
	} else {
		result = right
		result.next = sortedMerge(left, right.next)
	}
	return result
}

// Partition List: Given a value x, partition the list such that all nodes less than x come before nodes greater than or equal to x.
func (list *LinkedList) Partition(x int) {
	var lessHead, lessTail, greaterHead, greaterTail *Node

	current := list.head
	for current != nil {
		next := current.next
		current.next = nil
		if current.data < x {
			if lessHead == nil {
				lessHead = current
				lessTail = lessHead
			} else {
				lessTail.next = current
				lessTail = current
			}
		} else {
			if greaterHead == nil {
				greaterHead = current
				greaterTail = greaterHead
			} else {
				greaterTail.next = current
				greaterTail = current
			}
		}
		current = next

	}

	if lessHead == nil {
		list.head = greaterHead
	} else {
		list.head = lessHead
		lessTail.next = greaterHead
	}
}

// Rotate List: Implement a function to rotate the linked list to the right by k places.
/*
To rotate a singly linked list to the right by k places, we need to follow these steps:
Calculate the length of the list: Traverse the list to determine its length.
Connect the list into a circular list: Link the last node to the head of the list.
Find the new head: Calculate the new head position by considering the effective number of rotations needed (k % length).
*/
func (list *LinkedList) Length() int {
	length := 0
	current := list.head

	for current != nil {
		length++
		current = current.next
	}
	return length
}

// RotateRight rotates the linked list to the right by k places
func (list *LinkedList) RotateRight(k int) {
	if list.head == nil || k <= 0 {
		return
	}

	// calculate the lenght of the list
	lenght := list.Length()
	k = k % lenght // Effective rotation needed

	if k == 0 {
		return
	}

	// find the tail of the  list and make the list circular
	tail := list.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = list.head

	// Find the new tail: (length - k) steps from the beginning
	newTail := list.head
	for i := 0; i < lenght-k-1; i++ {
		newTail = newTail.next
	}

	// The new head is next of the new tail
	newHead := newTail.next
	newTail.next = nil

	// Update the head
	list.head = newHead
}

// Add Two Numbers: Write a function to add two numbers represented by linked lists.
/*
To add two numbers represented by linked lists, we can follow these steps:
Traverse both linked lists simultaneously: Add corresponding nodes of the two lists.
Handle carry: If the sum of two digits is greater than or equal to 10, carry the extra digit to the next addition.
Create a new linked list: Store the sum of the digits in a new linked list.
*/
func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList) InsertionAtEnd(data int) {
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

func AddTwoLists(list1, list2 *LinkedList) *LinkedList {
	result := &LinkedList{}
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
		result.InsertionAtEnd(sum % 10)

	}

	if carry > 0 {
		result.InsertionAtEnd(carry)
	}
	return result
}

// Swap Nodes in Pairs: Implement a function to swap every two adjacent nodes in a linked list.
func (list *LinkedList) SwapPairs() {
	dummy := &Node{next: list.head}
	prev := dummy
	current := list.head

	for current != nil && current.next != nil {
		first := current
		second := current.next

		// swapping nodes
		prev.next = second
		first.next = second.next
		second.next = first

		// Move pointers forward
		prev = first
		current = first.next
	}
	list.head = dummy.next
}

// Reorder List: Write a function to reorder the list such that it goes from L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → ...

// FindMiddle finds the middle of the linked list
/*func (list *LinkedList) FindMiddle() *Node {
	slow, fast := list.head, list.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}*/

// ReverseList reverse the  linked list and returns the new head
func ReverseList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

// MergeList merges two lists by alternating nodes
func MergeLists(list1, list2 *Node) {
	for list2 != nil {
		list1Next, list2Next := list1.next, list2.next
		list1.next = list2

		if list1.next == nil {
			break
		}
		list2.next = list1Next
		list1 = list1Next
		list2 = list2Next
	}
}

// ReorderList reorders the list as per the given pattern
func (list *LinkedList) ReorderList() {
	if list.head == nil || list.head.next == nil {
		return
	}

	// Find the middle of the list
	middle := list.FindMiddle()

	// split the list into two halves
	secondHalf := middle.next
	middle.next = nil

	// Reverse the second half
	secondHalf = ReverseList(secondHalf)

	// Merge the  two halves
	MergeLists(list.head, secondHalf)
}
