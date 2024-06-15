# Day 5: Circular Doubly Linked Lists

## Theory
A circular doubly linked list is a type of linked list where each node contains a reference to both the next and previous nodes, and the last node connects back to the first node, forming a circular structure.

### Structure of a Node

Each node in a circular doubly linked list typically has three components:

Data: Stores the data value.
Next: Pointer to the next node in the list.
Prev: Pointer to the previous node in the list.

## Benefits


### Bidirectional Circular Traversal:

You can traverse the list in both forward and backward directions starting from any node, making it flexible for various operations.
Useful in Complex Applications:
Music Playlists: Allows seamless looping of songs.
Circular Buffers: Useful in managing buffers in a circular manner.
Round-Robin Scheduling: Effective for algorithms that require cyclic scheduling.

## Operations Specific to Circular Doubly Linked Lists

### Insertion

At the Beginning:
Insert a new node before the current head. Adjust pointers to maintain the circular structure.

### At the End:

Insert a new node after the current tail. Adjust pointers so that the new node points back to the head, and the old tail points to the new node.
At a Specific Position:
Traverse to the desired position and insert the new node. Update the pointers of the adjacent nodes to include the new node.

### Deletion

### From the Beginning:

Remove the head node and update the head to the next node. Adjust the tail's next pointer to the new head and the new head's previous pointer to the tail.
From the End:
Remove the tail node and update the tail to the previous node. Adjust the new tail's next pointer to the head and the head's previous pointer to the new tail.
From a Specific Position:
Traverse to the node to be deleted. Adjust the pointers of the adjacent nodes to exclude the node being removed.
Traversal

### Forward Traversal:

Start from any node and follow the next pointers to visit each node in the list. Continue until the starting node is reached again.
Backward Traversal:
Start from any node and follow the previous pointers to visit each node in the list. Continue until the starting node is reached again.

### More Examples and Explanations

For more detailed explanations and code examples, you can visit the following URL: [Circular Doubly Linked List Operations](https://github.com/helloabhii/go-dsa/tree/master/Week-02/Day-05/Ques/main.go)

