# Day 1: Introduction to Stacks

## Concepts

### Definition and Characteristics of a Stack (LIFO - Last In, First Out):

Stack: A stack is an abstract data type that serves as a collection of elements, with two principal operations:
Push: Adds an element to the collection.
Pop: Removes the most recently added element that has not yet been removed.
LIFO: Last In, First Out describes the order in which elements are accessed. The last element added to the stack will be the first one to be removed.

### Use Cases:

Undo Functionality: Many software applications use stacks to implement the undo feature. Every time a user performs an action, that action is pushed onto a stack. When the user wants to undo an action, the most recent action is popped from the stack.
Expression Evaluation: Stacks are used to evaluate expressions, particularly those written in postfix notation (also known as Reverse Polish Notation).
Backtracking Algorithms: Stacks are used in various backtracking algorithms, such as those for solving mazes, navigating trees, and parsing expressions.

### Operations

#### Push (Add an element to the top):

##### Algorithm:

Check if the stack is full (if using a fixed-size stack).
Increment the top index.
Add the new element at the top index.
Complexity: O(1)
Pop (Remove the top element):

##### Algorithm:

Check if the stack is empty.
Retrieve the top element.
Decrement the top index.
Return the retrieved element.
Complexity: O(1)
Peek/Top (View the top element without removing it):

##### Algorithm:

Check if the stack is empty.
Return the element at the top index.
Complexity: O(1)
IsEmpty (Check if the stack is empty):

##### Algorithm:

Return true if the top index is -1.
Return false otherwise.
Complexity: O(1)


##### For more information:

