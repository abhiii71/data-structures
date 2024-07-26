# Day 1: Introduction to Hash Tables


## What is a Hash Table?

### Definition:

A hash table is a data structure that provides fast insertion, deletion, and retrieval of data using a hash function to compute an index into an array of buckets or slots.

### Advantages:

Average time complexity for insertion, deletion, and search operations is O(1).
Useful for implementing associative arrays, database indexing, caches, and more.

### Use Cases and Importance:

Hash Maps: Store key-value pairs where keys are hashed to determine the location in memory.
Caches: Quickly retrieve data that is accessed frequently.
Indexing: Speed up database queries by quickly locating records.

### Basic Terminology:

#### Hash Function:
A function that maps input keys to indices in the hash table.
Should minimize collisions and uniformly distribute entries across the table.

#### Hash Value:

The result of applying the hash function to a key, determining the index in the hash table.

#### Keys and Values:
Keys: Unique identifiers used to store and retrieve values.
Values: Data associated with each key.

Afternoon: Basic Operations

### Insertion Operation:

#### Process:
Apply the hash function to the key to get an index.
Insert the key-value pair into the array at the calculated index.

### Collision Handling:

Chaining: Use linked lists to handle collisions at the same index.
Open Addressing: Find another slot in the array based on a probing sequence.

### Deletion Operation:

#### Process:

Apply the hash function to the key to find the index.
Remove the key-value pair from the array.
Handle collisions appropriately to ensure the hash table remains functional.

### Search Operation:

#### Process:

Apply the hash function to the key to find the index.
Check if the key exists at the index or in the linked list (in chaining) or through probing (in open addressing).

### Understanding Load Factor:

Definition: The load factor is the ratio of the number of elements to the number of buckets in the hash table.

Formula: Load Factor = Number of Elements / Number of Buckets

#### Impact on Performance:

High Load Factor: More collisions, potentially increased lookup time.
Low Load Factor: Fewer collisions, faster operations.

### Dynamic Resizing:

Rehashing: Increase the number of buckets and redistribute existing elements to maintain efficient performance.
When to Rehash: Typically when the load factor exceeds a certain threshold.


