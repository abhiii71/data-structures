# Day 4: Advanced Sorting Techniques - Deep Dive

## Merge Sort

### Understanding Merge Sort Algorithm:

Concept: Merge Sort is a divide-and-conquer algorithm that divides the array into two halves, recursively sorts each half, and then merges the sorted halves.

### Process:

Divide the array into two halves.
Recursively sort the two halves.
Merge the two sorted halves to form a single sorted array.

### Implementation:

Merge Function: Merges two sorted subarrays into one sorted array.
Recursive Sort Function: Recursively divides and sorts the array.


### Time and Space Complexity Analysis:

Time Complexity: O(n log n) for all cases (best, average, worst) due to the recursive division and merging.
Space Complexity: O(n) due to the use of temporary arrays for merging.

## Quick Sort

### Understanding Quick Sort Algorithm:

Concept: Quick Sort is a divide-and-conquer algorithm that selects a 'pivot' element, partitions the array around the pivot, and recursively sorts the partitions.

### Process:
Choose a pivot element.
Partition the array such that elements less than the pivot are on the left and elements greater than the pivot are on the right.
Recursively sort the subarrays on the left and right of the pivot.

### Implementation:

Partition Function: Partitions the array around the pivot and returns the index of the pivot.
Recursive Sort Function: Recursively sorts the partitions.

### Pivot Selection Strategies and Partitioning:

#### Choosing Pivot:

Last element (as implemented above).
First element.
Random element.
Median of three (first, middle, last).

#### Partitioning:

Lomuto partition scheme (as implemented above).
Hoare partition scheme (uses two indices moving towards each other).

### Time and Space Complexity Analysis:


### Time Complexity:

Best/Average Case: O(n log n) when partitioning is balanced.
Worst Case: O(n^2) when partitioning is highly unbalanced (e.g., already sorted array with poor pivot choice).

### Space Complexity:

O(log n) for the recursive stack space.

