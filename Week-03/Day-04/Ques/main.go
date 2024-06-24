package main

import "fmt"

type MyQueue struct {
	inStack  []int
	outStack []int
}

// MovingAverage struct to store the size, window of elements, and the current sum
type MovingAverage struct {
	size   int
	window []int
	sum    int
}

func main() {
	// 1. Implement Queue using Stacks (LeetCode #232)
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.inStack)
	fmt.Println("Peek element: ", q.Peek())
	fmt.Println("Popping the first element:  ", q.Pop())
	fmt.Println(q.outStack)
	fmt.Println("Popping next element: ", q.Pop())
	fmt.Println("Is Queue empty: ", q.Empty())

	// 2. Moving Average from Data Stream (LeetCode #346)
	m := Constructorr(3)
	fmt.Println(m.Next(1))
	fmt.Println(m.Next(10))
	fmt.Println(m.Next(3))
	fmt.Println(m.Next(5))
}

/*
The problem is to implement a queue using two stacks.
A queue is a First-In-First-Out (FIFO) data structure, while a stack is a Last-In-First-Out (LIFO) data structure.
By using two stacks, we can simulate the behavior of a queue.

Solution in Go
The solution involves using two stacks:

inStack: used for enqueue operations (Pushing elements onto the queue).
outStack: used for dequeue operations (Popping elements from the queue).
*/

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) IsEmpty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	q.peek()

	val := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return val
}

func (q *MyQueue) Peek() int {
	q.peek()
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func (q *MyQueue) peek() {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			val := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			q.outStack = append(q.outStack, val)
		}
	}
}

/*
// 2. Moving Average from Data Stream (LeetCode #346)

In this problem, you are given a stream of integers and a window size.
The task is to calculate the moving average of all integers in the sliding window.
A moving average is the average of a fixed number of elements,
which changes as new elements are added and old elements are removed from the fixed-size window.


Let's say the window size n is 3, and the stream of integers is [1, 2, 3, 4, 5]. Here's how the moving average would be calculated:

First integer (1):

Window: [1]
Moving average: 1 / 1 = 1.0
Second integer (2):

Window: [1, 2]
Moving average: (1 + 2) / 2 = 1.5
Third integer (3):

Window: [1, 2, 3]
Moving average: (1 + 2 + 3) / 3 = 2.0
Fourth integer (4):

Window: [2, 3, 4] (1 is removed)
Moving average: (2 + 3 + 4) / 3 = 3.0
Fifth integer (5):

Window: [3, 4, 5] (2 is removed)
Moving average: (3 + 4 + 5) / 3 = 4.0


*/

// Constructor function to initialize a MovingAverage with a given size
func Constructorr(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}

// Next method to add a new  value and calculate the moving average
func (m *MovingAverage) Next(val int) float64 {
	if len(m.window) == m.size {
		// Remove the oldest element from the window and subtrack its value from the sum
		m.sum -= m.window[0]
		m.window = m.window[1:]
	}
	// Add the new value to the window and update the sum
	m.window = append(m.window, val)
	m.sum += val

	// Return the moving average
	return float64(m.sum) / float64(len(m.window))
}
