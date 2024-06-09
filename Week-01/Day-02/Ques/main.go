package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{22, 7, 21, 69, 34, 55, 23, 7, 71, 69}
	fmt.Println("Array: ", arr)

	// Traverse and Print:

	fmt.Println("Traversed Array: ")
	traverseArr(arr)

	// Search Element:
	search := 71
	fmt.Printf("Search Element = %d : ", search)
	result := searchElement(arr, search)
	if result != -1 {
		fmt.Println("Element found in the Array.")
	} else {
		fmt.Println("Element not found in the Array.")
	}

	// Count Occurrences:
	element := 71
	fmt.Printf("%d occurred in Array %d time.\n", element, countOccur(arr, element))

	// Replace Element:
	arr2 := []int{21, 22, 7, 4, 2, 4}
	old := 4
	new := 69
	fmt.Println(replaceOccur(arr2, old, new))

	// Sort Array:
	fmt.Print("Sorted Array: ")
	sortedArray := sortArr(arr)
	fmt.Println(sortedArray)

	// Find Duplicate:
	fmt.Print("Duplicate Array members:")
	duplicateArray := duplicateFind(arr)
	fmt.Println(duplicateArray)

	// Unique Elements:
	fmt.Print("Unique Array members:")
	uniqueArray := uniqueElement(arr)
	fmt.Println(uniqueArray)

	// Frequency Count:
	fmt.Print("Frequency Count: ")
	Fcount := freqCount(arr)
	fmt.Println(Fcount)

	// Left Rotate:
	arr = []int{22, 7, 21, 69, 34, 55, 23, 7, 71, 69}
	fmt.Print("Original Array: ", arr, "\n")
	fmt.Print("Left Rotate: ")
	lRotate := leftRotate(arr)
	fmt.Println(lRotate)

	// Right Rotate:
	arr = []int{22, 7, 21, 69, 34, 55, 23, 7, 71, 69}
	fmt.Print("Right Rotate: ")
	rRotate := rightRotate(arr)
	fmt.Println(rRotate)

	// Sum Pairs:
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 7
	pairs := sumPairs(arr, target)
	fmt.Println("Pairs with sum equal to target: ", pairs)

	// Intersection:
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 = []int{22, 7, 21, 69, 34, 55, 23, 7, 71, 69}
	intersection := interSection(arr1, arr2)
	fmt.Println("Intersection of the Array 1 and Array 2: ", intersection)

	// Union:
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 = []int{22, 7, 21, 69, 34, 55, 23, 7, 71, 69}
	union := union(arr1, arr2)
	fmt.Println("Union of the Array 1 and Array 2: ", union)

	// Max Subarray Sum:
	k := 3
	fmt.Println("Original Array: ", arr)
	fmt.Println("k: ", k)
	result, err := maxSubarray(arr, k)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Maximum sum of subarray of size", k, "is", result)
	}

	// Cyclic Rotation:

	cRotation := cyclicRotation(arr)
	fmt.Println("Array after cyclic Rotation", cRotation)
}

// Traverse and Print: Traverse an array and print each element.
func traverseArr(arr []int) {
	for index := range arr {
		fmt.Println(arr[index])
	}
}

// Search Element: Search for a specific element in an array.
func searchElement(arr []int, search int) int {
	for index := range arr {
		if arr[index] == search {
			return search
		}
	}
	return -1
}

// Count Occurrences: Count the occurrences of a specific element.
func countOccur(arr []int, element int) int {
	counter := 0
	for index := range arr {
		if arr[index] == element {
			counter++
			return counter
		}
	}
	return counter
}

// Replace Element: Replace all occurrences of an element with another element.
func replaceOccur(arr []int, old, new int) []int {
	for index := range arr {
		if arr[index] == old {
			arr[index] = new
		}
	}
	return arr
}

// Sort Array: Sort an array in ascending order.
func sortArr(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// Find Duplicates: Find all duplicate elements in an array.
func duplicateFind(arr []int) []int {
	countElement := make(map[int]int)
	duplicateArr := []int{}

	for _, element := range arr {
		countElement[element]++
	}

	for element, count := range countElement {
		if count > 1 {
			duplicateArr = append(duplicateArr, element)
		}
	}
	return duplicateArr
}

// Unique Elements: Print all unique elements in an array.
func uniqueElement(arr []int) []int {
	countElement := make(map[int]int)
	uniqueArr := []int{}

	for _, element := range arr {
		countElement[element]++
	}

	for element, count := range countElement {
		if count == 1 {
			uniqueArr = append(uniqueArr, element)
		}
	}
	return uniqueArr
}

// Frequency Count: Count the frequency of each element in an array.
func freqCount(arr []int) map[int]int {
	frequencyCount := make(map[int]int)

	for _, element := range arr {
		frequencyCount[element]++
	}

	return frequencyCount
}

// Left Rotate: Rotate the array to the left by one position.
func leftRotate(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left := arr[0]
	n := len(arr)
	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = left
	return arr
}

// Right Rotate: Rotate the array to the right by one position.
func rightRotate(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	right := arr[len(arr)-1]
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = right
	return arr
}

// Sum Pairs: Find all pairs of elements whose sum is equal to a given number.
func sumPairs(arr []int, target int) [][2]int {
	seen := make(map[int]bool)
	pairs := [][2]int{}

	for _, element := range arr {
		complement := target - element
		if seen[complement] {
			pairs = append(pairs, [2]int{element, complement})
		}
		seen[element] = true
	}
	return pairs
}

// Intersection: Find the intersection of two arrays.
func interSection(arr1, arr2 []int) []int {
	intersection := []int{}

	arrOne := make(map[int]bool)

	// Add elements of the first array to the map
	for _, element := range arr1 {
		arrOne[element] = true
	}

	// Check elements of the second array for interSection
	for _, element := range arr2 {
		if arrOne[element] {
			intersection = append(intersection, element)
			delete(arrOne, element)
		}
	}
	return intersection
}

// Union: Find the union of two arrays.
func union(arr1, arr2 []int) []int {
	union := []int{}

	arrOne := make(map[int]bool)

	for _, element := range arr1 {
		arrOne[element] = true
	}
	for _, element := range arr2 {
		arrOne[element] = true
	}

	for key := range arrOne {
		union = append(union, key)
	}
	return union
}

// Max Subarray Sum: Find the maximum sum of a subarray of size k.
func maxSubarray(arr []int, k int) (int, error) {
	n := len(arr)
	if n < k {
		return 0, errors.New("array length is less than k")
	}

	// Compute the window over the key
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}

	maxSum := currentSum

	// Slide the window over the array
	for i := k; i < n; i++ {
		currentSum = currentSum - arr[i-k] + arr[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum, nil
}

// Cyclic Rotation: Rotate an array cyclically by one.
func cyclicRotation(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

	lastElement := arr[n-1]

	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}

	arr[0] = lastElement

	return arr
}
