package main

import (
	"fmt"
)

// over array
type ArrayStack struct {
	elements []int // slice of integers that holds the stack's elements
}

// over linked list
type Node struct {
	value int
	next  *Node
}

type LinkedListStack struct {
	top *Node
}

// Min Stack (LeetCode #155)
type MinStack struct {
	stack    []int
	minStack []int
}

func main() {
	// Testing Array-based stack
	ArrayStack := &ArrayStack{}

	fmt.Println("Array Stack elements initially: ", ArrayStack.elements)
	// over array
	ArrayStack.Push(1)
	ArrayStack.Push(2)
	ArrayStack.Push(3)

	// Handle the two return values from Peek
	top, ok := ArrayStack.Peek()
	if ok {
		fmt.Println("Top element:", top)
	} else {
		fmt.Println("Stack is empty")
	}

	// Array stack elemets
	fmt.Println("Array stack elements after push: ", ArrayStack.elements)

	fmt.Println("Is stack empty?", ArrayStack.IsEmpty()) // false

	// Correctly handle the two return values from Pop
	popped, ok := ArrayStack.Pop()
	if ok {
		fmt.Println("Popped element:", popped)
	} else {
		fmt.Println("Stack is empty")
	}

	fmt.Println("Is stack empty?", ArrayStack.IsEmpty()) // false

	if popped, ok := ArrayStack.Pop(); ok {
		fmt.Println("Popped element:", popped)
	} else {
		fmt.Println("Stack is empty")
	}

	ArrayStack.Pop()
	fmt.Println("Is stack empty?", ArrayStack.IsEmpty()) // true

	// Testing Linkedlist based stack
	LinkedListStack := &LinkedListStack{}
	LinkedListStack.Push(1)
	LinkedListStack.Push(2)
	LinkedListStack.Push(3)

	if top, ok := LinkedListStack.Peek(); ok {
		fmt.Println("Linked list stack - Top element: ", top)
	} else {
		fmt.Println("Linked list stack - stack is empty")
	}

	fmt.Println("Linked List stack - Is stack empty? ", LinkedListStack.IsEmpty())

	if popped, ok := LinkedListStack.Pop(); ok {
		fmt.Println("Linked list stack - Popped element: ", popped)
	} else {
		fmt.Println("Linked list stack - stack is empty")
	}

	fmt.Println("Linked list stack  - Is stack empty? ", LinkedListStack.IsEmpty())

	// Valid Parentheses (LeetCode #20)
	fmt.Println(IsValid("()"))
	fmt.Println(IsValid("{}[]()"))
	fmt.Println(IsValid("(]"))

	//  Min Stack (LeetCode #155)
	minStack := constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println(minStack.GetMin())

	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

// Push Method:
// Adds an element to the top of the stack
func (s *ArrayStack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop Method:
// Removes and returns the top element of the stack
func (s *ArrayStack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, true
}

// Peek Method:
// Returns the top element of the stack without removing it
func (s *ArrayStack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	return s.elements[len(s.elements)-1], true
}

// IsEmpty Method:
// Checks if the stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return len(s.elements) == 0
}

// OVER LINKEDLIST

// Push add an element to the  top of the stack
func (s *LinkedListStack) Push(value int) {
	newNode := &Node{value: value}
	if s.top != nil {
		newNode.next = s.top
	}
	s.top = newNode
}

// Pop removes and returns the top element of the stack
func (s *LinkedListStack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	value := s.top.value
	s.top = s.top.next
	return value, true
}

// Peek returns the top element of the stack without removing it
func (s *LinkedListStack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.value, true
}

// Is empty checks if the stack is empty
func (s *LinkedListStack) IsEmpty() bool {
	return s.top == nil
}

// Valid Parentheses (LeetCode #20)
func IsValid(s string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// MIN STACK
func constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.minStack) == 0 || x <= s.GetMin() {
		s.minStack = append(s.minStack, x)
	}
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if top == s.GetMin() {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return -1 // or any other default value indicating empty stack
	}

	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.minStack) == 0 {
		return -1 // or any other default value indicating empty stack
	}
	return s.minStack[len(s.minStack)-1]
}
