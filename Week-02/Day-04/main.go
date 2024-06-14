package main

import "fmt"

// Node represents a node in a circular linked list
type Node struct {
	data int
	next *Node
}

// CircularLinkedList represents a circular linked list
type CircularLinkedList struct {
	head *Node
}

func main() {
	list := &CircularLinkedList{}
	list.InsertAtBeginning(69)
	list.InsertAtEnd(7)
	list.InsertAtEnd(71)
	list.InsertAtBeginning(5)

	fmt.Println("Circular Linked List: ")
	list.PrintList()

	list.DeleteNode(5)
	fmt.Println("List after deleting node with data 5: ")
	list.PrintList()
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (list *CircularLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
	} else {
		temp := list.head
		for temp.next != list.head {
			temp = temp.next
		}
		temp.next = newNode
		newNode.next = list.head
	}
}

// InsertAtEnd inserts a new node at the end of the list
func (list *CircularLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
	} else {
		temp := list.head
		for temp.next != list.head {
			temp = temp.next
		}
		temp.next = newNode
		newNode.next = list.head
	}
}

// DeleteNode deletes a node with specific data from the list
func (list *CircularLinkedList) DeleteNode(data int) {
	if list.head == nil {
		fmt.Println("List is empty bro")
		return
	}

	temp := list.head
	var prev *Node
	for temp.data != data {
		if temp.next == list.head {
			fmt.Println("Node not found")
			return
		}
		prev = temp
		temp = temp.next
	}
	if temp == list.head {
		prev = prev.next
		for prev.next != list.head {
			prev = prev.next
		}
		list.head = temp.next
		prev.next = list.head
	} else if temp.next == list.head {
		prev.next = list.head
	} else {
		prev.next = temp.next
	}
}

// PrintList prints all the nodes in the list
func (list *CircularLinkedList) PrintList() {
	if list.head == nil {
		fmt.Println("List is empty.")
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
	fmt.Println(temp.data) // Print the data of the last node pointing back to the head
}
