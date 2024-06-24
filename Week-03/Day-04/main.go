package main

import "fmt"

// Implementing a Queue using Arrays
type Queue struct {
	elements []int
}

// Implementing a Queue using a Linked List
type Node struct {
	val  int
	next *Node
}

type QueueL struct {
	front *Node
	rear  *Node
}

func main() {
	// Implementing a Queue using Arrays:
	queue := &Queue{}

	queue.Enqueue(71)
	queue.Enqueue(69)
	queue.Enqueue(22)
	val, _ := queue.Front()

	fmt.Println("Front Element: ", val)
	fmt.Println("Is queue empty: ", queue.IsEmpty())
	val, _ = queue.Dequeue()
	fmt.Println("Dequeued element: ", val)
	val, _ = queue.Dequeue()
	fmt.Println("Front Element: ", val)
	fmt.Println("Is queue empty: ", queue.IsEmpty())
	val, _ = queue.Dequeue()
	val, _ = queue.Dequeue()
	fmt.Println("Front Element: ", val)

	// Implementing a Queue using a Linked List:
	queueL := &QueueL{}
	queueL.Enqueue(21)
	queueL.Enqueue(71)
	queueL.Enqueue(69)

	val, _ = queueL.Front()
	fmt.Println("Front Element: ", val)
}

//  Implementing a Queue using Arrays:

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element of  the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}

	value := q.elements[0]
	q.elements = q.elements[1:]
	return value, true
}

// Front returns the front element without removing it
func (q *Queue) Front() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	return q.elements[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

//  Implementing a Queue using a Linked List

// Enqueue adds an element to the end of the queue
func (q *QueueL) Enqueue(val int) {
	newNode := &Node{val: val}
	if q.rear != nil {
		q.rear.next = newNode
	}

	q.rear = newNode
	if q.front == nil {
		q.front = newNode
	}
}

// Dequeuer removes and returns the front element of the queue
func (q *QueueL) Dequeue() (int, bool) {
	if q.front == nil {
		return 0, false
	}
	value := q.front.val
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return value, true
}

// Front returns the front element without removing it
func (q *QueueL) Front() (int, bool) {
	if q.front == nil {
		return 0, false
	}
	return q.front.val, true
}

// IsEmpty checks if the queue is empty
func (q *QueueL) IsEmpty() bool {
	return q.front == nil
}
