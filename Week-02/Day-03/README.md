# Day 3: Introduction to Doubly Linked Lists

## Theory
Structure and Benefits of Doubly Linked Lists
A doubly linked list is a type of linked list where each node contains three parts:

Data: The value stored in the node.
Next Pointer: A pointer/reference to the next node in the sequence.
Prev Pointer: A pointer/reference to the previous node in the sequence.

### Benefits:

Bidirectional Traversal: You can traverse the list in both forward and backward directions.
Easier Deletion: Deleting a node is easier since you have direct access to the previous node.


### 1. Reversing a Linked List

Reversing a linked list involves changing the direction of the links between nodes. Instead of pointing to the next node, each node will point to its previous node, effectively reversing the order of nodes in the list.

### 2. Detecting Loops in a Linked List

A loop in a linked list occurs when a node's next pointer points to one of the previous nodes, forming a cycle. This can cause infinite traversal. Floyd’s Cycle-Finding Algorithm, also known as the Tortoise and Hare algorithm, is commonly used to detect such loops.

### 3. Finding the Middle Element of a Linked List

Finding the middle element of a linked list can be done efficiently using two pointers, known as the slow and fast pointers. The slow pointer moves one step at a time, while the fast pointer moves two steps at a time. When the fast pointer reaches the end, the slow pointer will be at the middle of the list.

## Practice

### Reversing a Linked List

#### Steps:

1. Initialize three pointers: `prev` (initially `nil`), `current` (initially head of the list), and `next` (initially `nil`).
2. Traverse the list. For each node, store the next node in `next`, reverse the current node's pointer to point to `prev`, move `prev` to the current node, and move `current` to `next`.
3. Continue until `current` becomes `nil`.
4. Set the head of the list to `prev`.

### Detecting a Loop in a Linked List

#### Steps (Floyd’s Cycle-Finding Algorithm):

1. Initialize two pointers, `slow` and `fast`, both starting at the head of the list.
2. Move `slow` one step at a time and `fast` two steps at a time.
3. If there is a loop, `slow` and `fast` will meet at some point.
4. If `fast` reaches `nil`, there is no loop.

### Finding the Middle Element of a Linked List

#### Steps:

1. Initialize two pointers, `slow` and `fast`, both starting at the head of the list.
2. Move `slow` one step at a time and `fast` two steps at a time.
3. When `fast` reaches the end of the list, `slow` will be at the middle.

### More Examples and Explanations

For more detailed explanations and code examples, you can visit the following URL:
[Advanced Doubly Linked List Operations](https://github.com/helloabhii/go-dsa/tree/master/Week-02/Day-03/Ques/main.go)
