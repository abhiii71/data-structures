# Day 5: More Sorting Algorithms - Deep Dive

## Heap Sort

### Understanding Heap Sort Algorithm:

Concept: Heap Sort is a comparison-based sorting algorithm that uses a binary heap data structure.

### Process:

Build a max heap from the input data.
The largest item (root of the heap) is swapped with the last item of the heap.
Reduce the heap size by one and heapify the root.
Repeat until the heap size is reduced to 1.
Heap Construction and Heapify Operation:

Heap Construction: Convert the array into a max heap.
Heapify Operation: Ensure that a subtree with a given root node satisfies the heap property.

### Time and Space Complexity Analysis:

### Time Complexity:

Building the heap: O(n)
Heapify operations: O(n log n)
Space Complexity: O(1) (in-place sorting algorithm)


## Comparison and Analysis

### Performance Comparison:

### Merge Sort:

Time Complexity: O(n log n)
Space Complexity: O(n)
Stability: Stable

### Quick Sort:

Time Complexity: Average: O(n log n), Worst: O(n^2)
Space Complexity: O(log n) (recursive stack space)
Stability: Not Stable

### Heap Sort:

Time Complexity: O(n log n)
Space Complexity: O(1)
Stability: Not Stable

### Best Use Cases:

Merge Sort: Suitable for sorting linked lists, stable sorting requirements, and external sorting (large datasets that do not fit in memory).

Quick Sort: Generally faster for small to medium-sized arrays, and when in-place sorting is required.

Heap Sort: Useful for scenarios where constant space is a constraint, and a guaranteed O(n log n) time complexity is needed.

