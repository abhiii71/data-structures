package main

import "fmt"

type Node struct {
	data   int
	next   *Node
	child  *Node
	prev   *Node
	random *Node
}

type HeaderLinkedList struct {
	header *Node
}

func main() {
	// Insert at Beginning:
	list := NewHeaderLinkedList()
	list.InsertAtBeginning(71)
	list.InsertAtBeginning(21)
	list.InsertAtBeginning(69)
	list.PrintList()

	// Remove Nth Node from End:
	n := 2
	list.RemoveNthFromEnd(n)
	fmt.Printf("Removing %dnd element from End of the list: ", n)
	list.PrintList()

	// Intersection of Two Lists:
	list1 := NewHeaderLinkedList()
	list2 := NewHeaderLinkedList()

	list1.InsertAtBeginning(1)
	list1.InsertAtBeginning(2)
	list1.InsertAtBeginning(3)
	intersectingNode := &Node{data: 4}
	list1.header.next.next.next = intersectingNode
	list1.header.next.next.next.next = &Node{data: 5}

	list2.InsertAtBeginning(6)
	list2.InsertAtBeginning(7)
	list2.header.next.next = intersectingNode

	fmt.Println("List 1: ")
	list1.PrintList()

	fmt.Println("List 2: ")
	list2.PrintList()

	intersection := getInterSectionNode(list1, list2)
	if intersection != nil {
		fmt.Printf("Intersecting at node with value: %d\n", intersection.data)
	} else {
		fmt.Println("No intersection found.")
	}

	// Check Palindrome:
	if list.IsPalindrome() {
		fmt.Println("The list is a palindrome.")
	} else {
		fmt.Println("The list is  not a palindrome.")
	}

	// Partition List:
	list.InsertAtBeginning(3)
	list.InsertAtBeginning(5)
	list.InsertAtBeginning(8)
	list.InsertAtBeginning(5)
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(2)
	list.InsertAtBeginning(1)
	fmt.Print("Original list: ")
	list.PrintList()
	x := 5
	list.Partition(x)
	fmt.Printf("List after partitioning around %d: \n", x)
	list.PrintList()
	/* Something wrong with this one
	// Flatten a Multilevel Linked List:
	list4 := NewHeaderLinkedList()
	n1 := list4.InsertAtBeginning(1)
	n2 := list4.InsertAtBeginning(2)
	n3 := list4.InsertAtBeginning(3)
	n4 := list4.InsertAtBeginning(4)
	n5 := list4.InsertAtBeginning(5)

	// Adding child lists
	n2.child = &Node{data: 6, next: &Node{data: 7}}
	n4.child = &Node{data: 8, next: &Node{data: 9}}
	fmt.Println("Original List: ")
	list4.PrintList()

	list4.FlattenList()
	fmt.Println("Flattened List: ")
	list4.PrintList()

	*/

	// Rotate List:

	list.PrintList()
	list.RotateRight(2)

	// Reorder List:
	fmt.Print("Original List: ")
	list.PrintList()

	list.Reorder()
	fmt.Print("Reorderd List: ")
	list.PrintList()

	// Copy List with Random Pointer:

	list5 := &HeaderLinkedList{}
	node1 := list5.InsertAtBeginning(1)
	node2 := list5.InsertAtBeginning(2)
	node3 := list5.InsertAtBeginning(3)

	// setting up random pointers

	node1.random = node3
	node2.random = node1
	node3.random = node2

	fmt.Println("Original list with random pointers: ")
	list5.PrintWithRandom()

	copiedList := list5.Copy()

	fmt.Println("Copied list with random pointers: ")
	copiedList.PrintWithRandom()

	// Remove Zero Sum Subsequences:

	list3 := NewHeaderLinkedList()
	list3.InsertAtBeginning(1)
	list3.InsertAtBeginning(2)
	list3.InsertAtBeginning(-3)
	list3.InsertAtBeginning(3)
	list.InsertAtBeginning(1)
	list3.RemoveZeroSumSublists()
	fmt.Println("List after removing zero-sum sublists: ")
	list3.PrintList()
}

func NewHeaderLinkedList() *HeaderLinkedList {
	return &HeaderLinkedList{header: &Node{}}
}

// Insert at Beginning: Implement a function to insert a node at the beginning of the header linked list.
func (list *HeaderLinkedList) InsertAtBeginning(data int) *Node {
	newNode := &Node{data: data}
	newNode.next = list.header.next
	if list.header.next != nil {
		list.header.next.prev = newNode
	}
	list.header.next = newNode
	return newNode
}

func (list *HeaderLinkedList) PrintList() {
	current := list.header.next
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Write a function to remove the nth node from the end of a header linked list.
/*
To remove the nth node from the end of a header linked list, you can use a two-pointer approach.
This approach involves moving one pointer `n` steps ahead and then moving both pointers until the first pointer reaches the end.
This way, the second pointer will be at the node just before the one to be removed.
*/
func (list *HeaderLinkedList) RemoveNthFromEnd(n int) {
	if list.header == nil || list.header.next == nil {
		return
	}

	dummy := &Node{next: list.header.next}
	first := dummy
	second := dummy

	// Move first pointer n steps ahead
	for i := 0; i < n; i++ {
		if first.next == nil {
			return // n is greater than the length of the list
		}
		first = first.next
	}

	// Move both pointers until first reaches the end
	for first.next != nil {
		first = first.next
		second = second.next
	}

	// Remove the nth node from the  end
	second.next = second.next.next

	// Update the head in case the removed node was the first node
	list.header.next = dummy.next
}

// Write a function to find the intersection node of two header linked lists.
/*
To find the intersection node of two header linked lists, you can use the following approach:
Calculate the lengths of both linked lists.
Determine the difference in lengths.
Move the pointer of the longer list ahead by the difference in lengths.
Traverse both lists simultaneously until you find the intersection node.
*/

func getLength(list *HeaderLinkedList) int {
	length := 0
	current := list.header.next
	for current != nil {
		length++
		current = current.next

	}
	return length
}

func getInterSectionNode(list1, list2 *HeaderLinkedList) *Node {
	len1 := getLength(list1)
	len2 := getLength(list2)

	var longer, shorter *Node
	diff := 0

	if len1 > len2 {
		longer = list1.header.next
		shorter = list2.header.next
		diff = len1 - len2
	} else {
		longer = list2.header.next
		shorter = list1.header.next
		diff = len2 - len1
	}

	for diff > 0 {
		longer = longer.next
		diff--
	}

	for longer != nil && shorter != nil {
		if longer == shorter {
			return longer
		}
		longer = longer.next
		shorter = shorter.next
	}
	return nil
}

// Write a function to check if a header linked list is a palindrome.
/*
To check if a header linked list is a palindrome, you can use a two-pointer technique. Here's a step-by-step approach:
Find the middle of the linked list: Use the slow and fast pointer technique to find the middle of the linked list.
Reverse the second half of the linked list: Reverse the second half of the linked list.
Compare the first half and the reversed second half: Compare the nodes of the first half with the nodes of the reversed second half.
Restore the list (optional): Optionally, you can restore the list to its original state by reversing the second half again.
*/
func (list *HeaderLinkedList) IsPalindrome() bool {
	if list.header == nil || list.header.next == nil {
		return true
	}

	// Find the middle of the linked list

	slow, fast := list.header.next, list.header.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// Reverse the second half of the linked list
	var prev *Node
	current := slow
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	// Compare the first half and the reversed second half
	firstHalf, secondHalf := list.header.next, prev
	for secondHalf != nil {
		if firstHalf.data != secondHalf.data {
			return false
		}
		firstHalf = firstHalf.next
		secondHalf = secondHalf.next
	}
	return true
}

// Write a function to partition a header linked list around a value x, such that all nodes less than x come before nodes greater than or equal to x.
/*
Partitioning a linked list around a value `x` means rearranging the nodes in such a way that all nodes
with values less than `x` come before nodes with values greater than or equal to `x`.
The relative order of nodes in each partition should be preserved from the original list.

For example, if the linked list is `3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1` and `x` is `5`,
the partitioned list should be `3 -> 2 -> 1 -> 5 -> 8 -> 5 -> 10`
*/
func (list *HeaderLinkedList) Partition(x int) {
	if list.header == nil || list.header.next == nil {
		return
	}

	lessHead := &Node{}
	lessTail := lessHead
	greaterHead := &Node{}
	greaterTail := greaterHead

	current := list.header.next
	for current != nil {
		if current.data < x {
			lessTail.next = current
			lessTail = current
		} else {
			greaterTail.next = current
			greaterTail = current
		}
		current = current.next
	}
	// combine the two lists
	lessTail.next = greaterHead.next
	greaterTail.next = nil

	// Update the head of the original list
	list.header.next = lessHead.next
}

// Write a function to flatten a multilevel header linked list where each node can have a child linked list.
/*
To flatten a multilevel header linked list where each node can have a child linked list,
you need to traverse the list and recursively flatten each child list. Here's a step-by-step explanation and the code to achieve this:
Node Structure: Each node has a value, a pointer to the next node, and a pointer to a child node.
Flattening Process:
Traverse the main list.
If a node has a child, recursively flatten the child list.
Insert the flattened child list between the current node and the next node.
Continue the process until the entire list is flattened.
*/
/*
func (list *HeaderLinkedList) Flatten() {
	if list.header == nil {
		return
	}
	current := list.header.next
	for current != nil {
		if current.child != nil {
			// Flatten the child list
			childTail := current.child
			for childTail.next != nil {
				childTail = childTail.next
			}

			// Connect the tail to the next node
			childTail.next = current.next

// Connect the current node to the head of the child list
			current.next = current.child

			// Set the child pointer to nil
			current = current.next
		}
		current = current.next
	}
}
*/

// Write a function to rotate a header linked list to the right by k places.
/*
To rotate a header linked list to the right by `k` places, you need to follow these steps:
Determine the length of the linked list.
Compute the effective rotation, which is `k % length`.
Find the new tail and new head of the rotated list.
Adjust the pointers to rotate the list.
*/
func (list *HeaderLinkedList) RotateRight(k int) {
	if list.header == nil || list.header.next == nil || k == 0 {
		return
	}

	// Step 1: Determine the length of the list
	length := 0
	current := list.header.next
	for current != nil {
		length++
		current = current.next
	}
	// Step 2: Compute the effective roatation
	k = k % length
	if k == 0 {
		return
	}

	// Step 3: Find the new tail and new head
	slow, fast := list.header.next, list.header.next
	for i := 0; i < k; i++ {
		fast = fast.next
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	// Step 4: Adjust the pointers to rotate the list
	newHead := slow.next
	slow.next = nil
	fast.next = list.header.next
	list.header.next = newHead
}

// Write a function to copy a header linked list where each node has an additional random pointer.
func (list *HeaderLinkedList) PrintWithRandom() {
	if list.header == nil {
		return
	}
	current := list.header.next

	for current != nil {
		randomValue := "nil"
		if current.random != nil {
			randomValue = fmt.Sprintf("%d", current.random.data)
		}
		fmt.Printf("Node value: %d, Random points to: %s\n ", current.data, randomValue)
		current = current.next
	}
}

func (list *HeaderLinkedList) Copy() *HeaderLinkedList {
	if list.header == nil {
		return nil
	}

	// Step 1: Create a copy of each node and insert it right next to the original node
	current := list.header.next
	for current != nil {
		copyNode := &Node{data: current.data}
		copyNode.next = current.next
		current.next = copyNode
		current = copyNode.next
	}

	// Step 2: Set the random pointers for the copied nodes
	current = list.header.next
	for current != nil {
		if current.random != nil {
			current.next.random = current.random.next
		}
		current = current.next.next
	}

	// Step 3: Separate the copied nodes to form the new list
	newList := &HeaderLinkedList{header: &Node{}}
	current = list.header.next
	copyCurrent := newList.header
	for current != nil {
		copyNode := current.next
		current.next = copyNode.next
		copyCurrent.next = copyNode
		copyCurrent = copyNode
		current = current.next
	}
	return newList
}

// Write a function to reorder a header linked list such that the nodes are rearranged in a specific pattern (e.g., L0→Ln→L1→Ln-1→L2→Ln-2→…).
/*
Reordering a linked list in the pattern
`L0→Ln→L1→Ln-1→L2→Ln-2→…` means rearranging the nodes such that the first node is followed by the last node,
then the second node is followed by the second last node, and so on.
For example, if the linked list is `1 -> 2 -> 3 -> 4 -> 5`,
the reordered list should be `1 -> 5 -> 2 -> 4 -> 3`.
*/
func (list *HeaderLinkedList) Reorder() {
	if list.header == nil || list.header.next == nil {
		return
	}

	// Step 1: Find the middle of the list
	slow, fast := list.header.next, list.header.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// Step 2: Reverse the second half of the list
	var prev *Node
	current := slow
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	// Step 3: Merge the two halves
	first, second := list.header.next, prev
	for second.next != nil {
		temp1 := first.next
		temp2 := second.next

		first.next = second
		second.next = temp1

		first = temp1
		second = temp2
	}
}

//10 Write a function to remove consecutive nodes that sum up to zero from a header linked list.
/*
To remove consecutive nodes that sum up to zero from a header linked list,
we can use a hashmap (or dictionary) to keep track of the cumulative sums as we traverse the list.
The idea is to use the cumulative sum to identify sublists that sum to zero.

Initialize a dummy node to handle edge cases.
Traverse the list while maintaining the cumulative sum.
Use a hashmap to store the cumulative sum and the corresponding node.
If the cumulative sum is found in the hashmap,
remove the nodes between the previous occurrence and the current node.
Update the next pointers to skip the zero-sum sublist.
*/
func (list *HeaderLinkedList) RemoveZeroSumSublists() {
	if list.header == nil || list.header.next == nil {
		return
	}

	dummy := &Node{next: list.header.next}
	cumulativeSum := 0
	sumMap := make(map[int]*Node)
	sumMap[0] = dummy

	current := list.header.next
	for current != nil {
		cumulativeSum += current.data
		if prevNode, found := sumMap[cumulativeSum]; found {
			// Remove nodes between prevNode and current
			temp := prevNode.next
			sum := cumulativeSum
			for temp != current {
				sum += temp.data
				delete(sumMap, sum)
				temp = temp.next
			}
			prevNode.next = current.next
		} else {
			sumMap[cumulativeSum] = current
		}
		current = current.next
	}

	list.header.next = dummy.next
}
