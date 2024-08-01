# Day 6: Applications of Hash Tables

## What is Caching?

Caching is the process of storing copies of data in a high-speed storage layer (cache) to quickly access frequently used data. It improves the speed of data retrieval and reduces the need to access slower storage layers.

### Why Caching is Important:

Speed: By keeping frequently accessed data in a fast storage layer, the system can respond more quickly to requests.
Reduced Load: Caching reduces the number of requests to slower storage layers, decreasing latency and improving overall system performance.
Scalability: Efficient caching strategies can help systems scale by managing load more effectively.

### Types of Caches

#### Memory Cache:

Description: Stores data in RAM (Random Access Memory) for very fast access.
Use Case: Frequently accessed data that needs to be retrieved quickly, such as web page elements, database query results, and session data.
Advantages: Extremely fast read/write times.
Disadvantages: Limited by the size of RAM and volatility (data is lost on restart).

#### Disk Cache:

Description: Stores frequently accessed disk data to reduce disk I/O operations.
Use Case: File system data, large data sets, and data that does not fit entirely in memory.
Advantages: Larger storage capacity compared to memory cache.
Disadvantages: Slower than memory cache but still faster than accessing the original data source.

### Cache Strategies
#### 1. LRU (Least Recently Used):

Description: Evicts the least recently used items first.
Algorithm: Maintains a record of access order and removes the oldest accessed item when the cache is full.
Advantages: Simple and effective for many use cases where recently accessed data is likely to be accessed again.
Disadvantages: Can be less effective if the access pattern frequently changes.

#### 2. LFU (Least Frequently Used):

Description: Evicts the least frequently used items first.
Algorithm: Maintains a count of how often each item is accessed and removes the least accessed item when the cache is full.
Advantages: Works well for scenarios where some items are accessed much more frequently than others.
Disadvantages: Can be more complex to implement and may not perform well if there are sudden spikes in access patterns.

### Deep Dive into LRU and LFU

#### LRU (Least Recently Used) Cache Implementation

##### Concept:

Use a combination of a hash table and a doubly linked list.
The hash table provides O(1) time complexity for lookups.
The doubly linked list maintains the order of access with the most recently accessed item at the front and the least recently accessed item at the back.
Pseudo Code:

##### Initialization:

Create a hash table for storing key-node pairs.
Create a doubly linked list for maintaining the order of access.

##### Access/Insert Item:

If the item is already in the cache, move it to the front of the list.
If the item is not in the cache, add it to the front of the list and update the hash table.
If the cache is full, remove the item at the back of the list and delete its entry from the hash table.

##### Eviction:

Remove the item at the back of the list (least recently used).
