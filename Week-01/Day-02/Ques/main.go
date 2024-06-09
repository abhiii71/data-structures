package main

import (
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
func interSection(arr []int, arr []int) []int {
	return arr
}
