# Day 2: Array Traversal and Manipulation

## Theory

### Linear Search:

Linear search is the simplest search algorithm. 
It checks each element of the array sequentially until the target element is found or the end of the array is reached.

Time Complexity: O(n)

### Binary Search:
Binary Search is a more efficient algorithm but requires the array to be sorted. It repeatedly divides the search interval in half. 
If the target value is less than the middle element, it narrows the interval to the lower half. Otherwise, it narrows it to the upper half.

Time Complexity: O(log n)

### Bubble Sort:
Bubble sort repeatedly steps through the list, compares adjacent elements, 
and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted.

Time Complexity: O(n^2)

### Selection Sort: 

Selection sort divides the array into a sorted and an unsorted region. 
It repeatedly selects the smallest (or largest) element from the unsorted region 
and moves it to the end of the sorted region.

Time Complexity: O(n^2)

### Insertion Sort:  
Insertion sort builds the final sorted array one item at a time. 
It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

Time Complexity: O(n^2)

## Summary
