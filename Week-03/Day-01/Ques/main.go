package main

import "fmt"

// Implement Stack using Arrays
type ArrayStack struct {
	element []int
}

// Implement Stack using Linked List
type LinkedListStack struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func main() {
	// Over Array
	arr := &ArrayStack{}

	fmt.Println("Operations over Array: ")
	// Push Operation:
	arr.Push(1)
	arr.Push(2)

	// Pop Operation:
	val, status := arr.Pop()
	if status {
		fmt.Println("Element successfully popped out from stack: ", val)
	} else {
		fmt.Println("Stack is empty")
	}

	// Peek Operation:
	value, ok := arr.Peek()
	if ok {
		fmt.Println("Peek element: ", value)
	} else {
		fmt.Println("Stack is empty")
	}

	// IsEmpty Operation:
	empty := arr.IsEmpty()
	fmt.Println("Is Array Stack Empty: ", empty)

	fmt.Println("Elements in Array Stack: ", arr.element)
	//////////////////////////////////////////////////////////////////
	// over LinkedList
	list := &LinkedListStack{}

	fmt.Println("Operations over List: ")
	// Push Operation:
	list.Push(11)
	list.Push(12)
	list.Push(13)

	// Pop Operation:
	val, status = list.Pop()
	if status {
		fmt.Println("Popped out element from linkedlist stack: ", val)
	} else {
		fmt.Println("Stack is empty")
	}

	// Peek Operation:
	value, ok = list.Peek()
	if ok {
		fmt.Println("Peek element: ", value)
	} else {
		fmt.Println("stack is empty")
	}

	// IsEmpty Operation:
	empty = list.IsEmpty()
	fmt.Println("Is Linkedlist Stack Empty: ", empty)

	// Valid Parenthesis:
	fmt.Println("Checking Valid Parenthesis: ")
	input := "()[]{}"
	fmt.Printf("Input: %s\nOutput: %v\n", input, isValid(input))

	input = "([)]"
	fmt.Printf("Input: %s\nOutput: %v\n", input, isValid(input))

	input = "{[]}"
	fmt.Printf("Input: %s\nOutput: %v\n", input, isValid(input))
}

// Implement the Push operation for a stack.
func (s *ArrayStack) Push(value int) {
	s.element = append(s.element, value)
}

func (s *LinkedListStack) Push(value int) {
	newNode := &Node{value: value}
	newNode.next = s.head
	s.head = newNode
}

// Implement the Pop operation for a stack.
func (s *ArrayStack) Pop() (int, bool) {
	if s.element == nil {
		return 0, false
	}
	index := len(s.element) - 1
	value := s.element[index]
	s.element = s.element[:index]

	return value, true
}

func (s *LinkedListStack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	val := s.head.value
	s.head = s.head.next

	return val, true
}

// Implement the Peek operation to view the top element without removing it.
func (s *ArrayStack) Peek() (int, bool) {
	if s.element == nil {
		return 0, false
	}
	return s.element[len(s.element)-1], true
}

func (s *LinkedListStack) Peek() (int, bool) {
	if s.head == nil {
		return 0, false
	}

	return s.head.value, true
}

// Implement the IsEmpty operation to check if the stack is empty.
func (s *ArrayStack) IsEmpty() bool {
	return len(s.element) == 0
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.head == nil
}

/*

Valid Parentheses (LeetCode #20)

Given a string containing just the characters '(', ')', '{', '}', '[', ']', determine if the input string is valid.
Example:
plaintext
Copy code
Input: "()[]{}"
Output: true
Implement a function in Go to solve this problem using a stack.

*/

func isValid(s string) bool {
	stack := []rune{}
	matchingBracket := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != matchingBracket[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
