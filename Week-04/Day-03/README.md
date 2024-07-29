# Day 3: Collision Resolution Techniques - Chaining
Morning: Conceptual Understanding
## What are Collisions and Why Do They Occur?

### Definition:

A collision occurs when two different keys hash to the same index in a hash table.
This results in multiple entries mapping to the same location, causing a conflict.

### Causes:

#### Limited Size of the Hash Table:

When the hash table has fewer slots than the number of keys, some keys will inevitably hash to the same index.

#### Poorly Designed Hash Functions:

Hash functions that do not distribute keys uniformly across the table increase the likelihood of collisions.
Chaining as a Collision Resolution Technique

### Definition:

Chaining resolves collisions by maintaining a list of all elements that hash to the same index.

#### Process:

Each slot in the hash table points to a linked list of entries that map to the same hash value.
When inserting a key-value pair, if a collision occurs, the new entry is appended to the linked list at the corresponding index.
During search, deletion, or update operations, the linked list at the hashed index is traversed to find the desired entry.

## Advantages and Disadvantages of Chaining

### Advantages:

#### Simple to Implement:
The concept of linked lists is straightforward and easy to understand.

#### Can Handle an Unlimited Number of Collisions (Subject to Memory):

Linked lists can grow dynamically to accommodate any number of collisions.

#### Efficient Average Case Time Complexity (O(1)):

Inserting, deleting, and searching operations can be done in constant time, assuming uniform distribution of keys.

### Disadvantages:

#### Extra Memory Required for Pointers in Linked Lists:

Each element in the linked list requires additional memory for pointers.

#### Performance Degrades with Long Linked Lists:

If many keys hash to the same index, the linked list becomes long, and operations degrade to O(n) in the worst case, where n is the number of elements in the list.

