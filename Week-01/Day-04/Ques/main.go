package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{71, 21, 2, 69, 22, 9}
	fmt.Printf("Array: %d\n", arr)

	// Bubble Sort:
	fmt.Printf("Bubble Sort: %d  \n", bubbleSort(arr))

	// Selection Sort:
	fmt.Printf("Selection Sort: %d\n", selectionSort(arr))

	// Insertion Sort:
	fmt.Printf("Insertion Sort: %d\n", insertionSort(arr))

	// Sort Descending:
	fmt.Printf("Descending Sort: %d\n", descBubbleSort(arr))

	// Count Swaps in Bubble Sort:
	fmt.Printf("Count Swap in Bubble Sort: %d\n", bSortCount(arr))

	// Sort Strings:
	arr1 := []string{"banana", "apple", "cherry", "mango", "blueberry"}
	fmt.Printf("Sort String Using Insertion Sort: %s\n", strInsertionSort(arr1))

	// Check Sorted:
	arr = []int{71, 21, 2, 69, 22, 9}
	fmt.Println("Original Array: ", arr)
	check := checkSort(arr)
	fmt.Printf("Checking Array is Sorted or not: %t\n", check)

	// Sort Subarray:
	arr = []int{64, 25, 12, 22, 11, 34, 90, 33}
	start, end := 2, 5
	sortSubArr := sortSubarray(arr, start, end)
	fmt.Printf("Sorted Subarray: %d\n", sortSubArr)

	// Stable Sort Check:
	StableSort := stableSort(arr)
	fmt.Printf("Stable Sort: %d\n", StableSort)

	// Demonstration of Stability
	arr3 := []Pair{
		{4, "a"},
		{3, "b"},
		{3, "c"},
		{1, "d"},
		{4, "e"},
		{3, "f"},
	}
	insertionSortPair(arr3)
	fmt.Println("Sorted array:", arr3)

	// ￼Find Median:
	arr = []int{1, 3, 5, 2, 4}
	bubsort(arr)
	fmt.Println(arr)
	median := findMedian(arr)
	fmt.Println("Median: ", median)

	// Bucket Sort:
	arr4 := []float64{0.42, 0.32, 0.33, 0.52, 0.37, 0.47, 0.51}
	array := bucketSort(arr4)
	fmt.Println("Sorted Array: ", array)

	// Sort by Frequency:
	sortedArr := sortByFrequency(arr)
	fmt.Println("Sort By Frequency: ", sortedArr)

	// Sort Negative and Positive:
	arr = []int{3, -2, 5, -1, 0, 4, -3}
	bSortNew(arr)
	fmt.Println("Sort Negative and Positive Array: ", arr)
}

// Bubble Sort: Implement Bubble Sort.
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Selection Sort: Implement Selection Sort.
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = arr[j]
			}
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}

	return arr
}

// Insertion Sort: Implement Insertion Sort.
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

// Sort Descending: Modify Bubble Sort to sort in descending order.
func descBubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Count Swaps in Bubble Sort: Count the number of swaps needed in Bubble Sort.
func bSortCount(arr []int) int {
	n := len(arr)
	swapCount := 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapCount++
			}
		}
	}
	return swapCount
}

// Sort Strings: Sort an array of strings using Insertion Sort.
func strInsertionSort(arr []string) []string {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1

		}
		arr[j+1] = key
	}
	return arr
}

// Check Sorted: Check if an array is already sorted.
func checkSort(arr []int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// Sort Subarray: Sort a specific subarray using Selection Sort.
func sortSubarray(arr []int, start, end int) []int {
	for i := start; i < end; i++ {
		minIdx := i
		for j := i + 1; j <= end; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// Stable Sort Check: Explain and demonstrate a stable sort.
/*
A stable sort is a sorting algorithm that preserves the relative order of records with equal keys (i.e., values).
In other words, if two elements are equal, they retain their original order in the sorted array.
Stability is important when the order of equal elements carries meaning,
such as when sorting by multiple criteria.
*/
func stableSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

// Demonstration of Stability:
type Pair struct {
	Value int
	Label string
}

func insertionSortPair(arr []Pair) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j].Value > key.Value {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// Adaptive Sort Check: Explain and demonstrate an adaptive sort.
/*
An adaptive sort is a sorting algorithm that takes advantage of existing order in the input data to improve its performance.
These algorithms are particularly efficient when sorting partially sorted or nearly sorted data, as they adapt their behavior based on the input characteristics.

Adaptive sorts perform well when the input data is already partially sorted or nearly sorted.

*/

// Sort Partially Sorted Array: Use Insertion Sort to sort a partially sorted array.

// Find Median: Find the median of an array using Bubble Sort.
func bubsort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func findMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2
	}
	return float64(arr[n/2])
}

// Bucket Sort Concept: Explain and demonstrate the concept of Bucket Sort.
/*
Bucket Sort is a sorting algorithm that works by distributing the elements of an array
into a number of buckets and then sorting each bucket individually,
typically using a different sorting algorithm, such as Insertion Sort or Quick Sort.
It is a distribution sort, and is useful when the input is uniformly distributed over a range.
*/
func bucketSort(arr []float64) []float64 {
	// number of buckets
	numBuckets := len(arr)

	// create empty buckets
	buckets := make([][]float64, numBuckets)
	for i := 0; i < numBuckets; i++ {
		buckets[i] = make([]float64, 0)
	}

	// Distribute elements into buckets
	for _, num := range arr {
		bucketIndex := int(num * float64(numBuckets))
		buckets[bucketIndex] = append(buckets[bucketIndex], num)
	}

	// sort individual buckets using Insertion Sort
	for i := 0; i < numBuckets; i++ {
		sort.Float64s(buckets[i])
	}

	// concatenate sorted buckets to produce the sorted array
	index := 0
	for i := 0; i < numBuckets; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}
	return arr
}

// Sort by Frequency: Sort elements by their frequency.

// Struct to store element and its frequency
type ElementFrequency struct {
	Element    int
	Frequency  int
	FirstIndex int
}

// Function to sort array by frequency
func sortByFrequency(arr []int) []int {
	// Step 1: Count frequencies and store first appearance index
	frequencyMap := make(map[int]int)
	indexMap := make(map[int]int)
	for i, num := range arr {
		frequencyMap[num]++
		if _, exists := indexMap[num]; !exists {
			indexMap[num] = i
		}
	}

	// Step 2: Store elements with frequencies and first appearance index in a slice
	elementFrequencies := make([]ElementFrequency, 0, len(frequencyMap))
	for element, frequency := range frequencyMap {
		elementFrequencies = append(elementFrequencies, ElementFrequency{
			Element:    element,
			Frequency:  frequency,
			FirstIndex: indexMap[element],
		})
	}

	// Step 3: Sort the slice by frequency, then by first appearance index
	sort.SliceStable(elementFrequencies, func(i, j int) bool {
		if elementFrequencies[i].Frequency == elementFrequencies[j].Frequency {
			return elementFrequencies[i].FirstIndex < elementFrequencies[j].FirstIndex
		}
		return elementFrequencies[i].Frequency > elementFrequencies[j].Frequency
	})

	// Step 4: Reconstruct the sorted array based on sorted frequencies
	sortedArr := make([]int, 0, len(arr))
	for _, ef := range elementFrequencies {
		for i := 0; i < ef.Frequency; i++ {
			sortedArr = append(sortedArr, ef.Element)
		}
	}

	return sortedArr
}

// Sort Negative and Positive: Sort an array with negative and positive numbers using Bubble Sort.
func bSortNew(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
