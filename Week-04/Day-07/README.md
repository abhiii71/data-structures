# Day 7: Assessment and Review

## 1. What is the average time complexity for search operations in a hash table?

A. O(n)
B. O(log n)
C. O(1)
D. O(n log n)
### Answer: C. O(1)

## Which collision resolution technique involves storing all elements that hash to the same index in a linked list?
A. Open Addressing
B. Chaining
C. Linear Probing
D. Double Hashing
### Answer: B. Chaining

## In open addressing, which technique checks the next available slot in a predefined sequence until an empty slot is found?

A. Linear Probing
B. Quadratic Probing
C. Double Hashing
D. All of the above
### Answer: D. All of the above

## In open addressing, which technique checks the next available slot in a predefined sequence until an empty slot is found?

A. Linear Probing
B. Quadratic Probing
C. Double Hashing
D. All of the above
### Answer: D. All of the above

### 1. What is a hash table?

Answer: A hash table is a data structure that maps keys to values for efficient lookup. It uses a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

### 2. What is a hash function?

Answer: A hash function is a function that takes an input (or key) and returns an integer, which is typically used as an index in a hash table. The aim is to distribute keys uniformly across the hash table to minimize collisions.

### 3. What are collisions in a hash table and how can they be resolved?

Answer: Collisions occur when two keys hash to the same index in a hash table. They can be resolved using techniques such as chaining (using linked lists) and open addressing (using probing sequences like linear probing, quadratic probing, and double hashing).

### 4. What is chaining in hash tables?

Answer: Chaining is a collision resolution technique where each bucket in the hash table points to a linked list of entries that hash to the same index. When a collision occurs, the new entry is added to the list.

### 5. What is open addressing in hash tables?

Answer: Open addressing is a collision resolution technique where, upon collision, the algorithm probes or searches for the next available slot according to a specified sequence (e.g., linear probing, quadratic probing, or double hashing).

### 6. What is the load factor in a hash table?

Answer: The load factor is a measure of how full the hash table is. It is calculated as the number of elements in the table divided by the number of slots (buckets). A high load factor may lead to more collisions and thus longer search times.

### 7. What is dynamic resizing in a hash table and why is it important?

Answer: Dynamic resizing involves expanding or shrinking the hash table when the load factor exceeds a certain threshold or becomes too low. This helps maintain efficient operations by reducing collisions or saving memory.

### 8. What is the average time complexity for insert, delete, and search operations in a hash table?

Answer: The average time complexity for insert, delete, and search operations in a hash table is O(1). However, in the worst case (e.g., all keys collide), the time complexity can degrade to O(n).

### 9. Explain the difference between linear probing and quadratic probing.

Answer: Both are open addressing techniques to resolve collisions. In linear probing, the next slot is checked sequentially (i.e., index + 1). In quadratic probing, the interval between probes increases quadratically (i.e., index + 1^2, index + 2^2, etc.), which helps avoid clustering.


