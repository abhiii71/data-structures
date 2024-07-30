# Day 4: Collision Resolution Techniques - Open Addressing

Open Addressing as a Collision Resolution Technique

#### Definition:

Open addressing handles collisions by finding another slot within the hash table for the colliding element. All elements are stored directly in the hash table array.

#### Different Strategies

### Linear Probing

#### Definition:

When a collision occurs, the algorithm checks the next slot in the array (i.e., index + 1). If the next slot is occupied, it continues checking subsequent slots until an empty slot is found.

#### Advantages:

Simple to implement.
Easy to understand.
Disadvantages:

Primary clustering: Continuous blocks of occupied slots can form, leading to longer search times

### Quadratic Probing

#### Definition:

When a collision occurs, the algorithm checks slots at intervals of 1, 4, 9, 16, etc. (i.e., index + 1^2, index + 2^2, index + 3^2, etc.). This helps in reducing clustering issues compared to linear probing.

#### Advantages:

Reduces primary clustering.
More evenly distributes keys.

#### Disadvantages:

Secondary clustering: Different keys can follow the same probing sequence.
More complex to implement.

### Double Hashing

#### Definition:

When a collision occurs, a secondary hash function is used to determine the interval for checking subsequent slots. The new slot is calculated using a formula: index + i * hash2(key), where i is the number of attempts.

### Advantages:

Minimizes clustering (both primary and secondary).
Provides a better distribution of keys.
Disadvantages:

More complex to implement.
Requires two good hash functions.


### Linear Probing

Linear probing is straightforward. When a collision occurs, it checks the next slot sequentially until an empty slot is found. This method is simple and easy to implement but can lead to primary clustering, where clusters of filled slots form, making future insertions slower.

#### Implementation Steps:

Compute the initial index using the primary hash function.
If the slot is empty, insert the key.
If the slot is occupied, move to the next slot (i.e., increment the index by 1) and repeat step 2.


### Quadratic Probing

Quadratic probing reduces primary clustering by checking slots at increasing intervals. Instead of moving to the next slot, it moves to slots at intervals of i^2.

#### Implementation Steps:

Compute the initial index using the primary hash function.
If the slot is empty, insert the key.
If the slot is occupied, calculate the next slot using i^2 (i.e., increment the index by i^2).
Repeat steps 2-3 until an empty slot is found.


### Double Hashing

Double hashing uses a second hash function to calculate the interval for probing. This reduces clustering more effectively than linear and quadratic probing.

#### Implementation Steps:

Compute the initial index using the primary hash function.
Compute the step size using the secondary hash function.
If the slot is empty, insert the key.
If the slot is occupied, move to the next slot using the formula index + i * stepSize.
Repeat steps 3-4 until an empty slot is found.


### Summary:
Open addressing is a collision resolution technique where all elements are stored directly in the hash table. It handles collisions by finding another slot within the hash table for the colliding element. Different strategies include linear probing, quadratic probing, and double hashing. Each has its own advantages and disadvantages, and understanding these can help in choosing the right method for a given application.
