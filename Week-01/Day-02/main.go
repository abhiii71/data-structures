package main

import "fmt"

func main() {
	// Traversing Array
	arr := []int{71, 12, 2, 4, 5, 69}
	fmt.Println("Traversing Array: ")
	traverse(arr)

	// Finding Element
	find := 71
	result := findElement(arr, find)
	if result != -1 {
		fmt.Println("Element Found At Index: ", findElement(arr, find))
	} else {
		fmt.Println("Element not Present in Array :)")
	}

	// Modifying Element
	fmt.Println("Modified Array: ", modifyElement(arr, 2, 71))

	fmt.Println("Replacing Element in Array: ", replaceElement(arr, 71, 69))
}

func traverse(arr []int) {
	for _, element := range arr {
		fmt.Println(element)
	}
}

func findElement(arr []int, find int) int {
	for index, value := range arr {
		if value == find {
			return index
		}
	}
	return -1
}

func modifyElement(arr []int, idx, new int) []int {
	arr[idx] = new
	return arr
}

func replaceElement(arr []int, old, new int) []int {
	for index := range arr {
		if arr[index] == old {
			arr[index] = new
		}
	}
	return arr
}
