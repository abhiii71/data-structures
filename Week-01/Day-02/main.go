package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	traverse()

	//finding element
	fmt.Println(find(arr, 5))

	//Modified
	mod(arr)

	// replace element
	fmt.Println("Replacing Element: ")
	fmt.Println(replace(arr, 1, 11))

	fmt.Println("Questions Started: ")

	// QUESTIONS:
	traversal(arr)

	fmt.Println(find_E(arr, 1))

	Modifying_E(arr)

	fmt.Println(replace_E(arr, 5, 71))

	fmt.Println(EdgeCase(arr, 69))
}

// Traversing array
func traverse() {
	arr := [5]int{1, 2, 3, 4, 5}
	for _, element := range arr {
		fmt.Println(element)
	}

}

// finding element
func find(arr []int, target int) int {
	for index, element := range arr {
		if element == target {
			return index
		}
	}
	return -1
}

//modifying element

func mod(arr []int) {
	arr[2] = 71
	fmt.Println(arr)
}

func replace(arr []int, old, new int) []int {
	for index, element := range arr {
		if arr[index] == old {
			arr[element] = new
		}
	}
	return arr
}

// QUESTIONS

func traversal(arr []int) {
	fmt.Println("Traversing Elements: ")
	for _, element := range arr {
		fmt.Println(element)
	}
}

// Finding Elements:
func find_E(arr []int, target int) int {
	for index, element := range arr {
		if element == target {
			fmt.Print("Element found at index: ")
			return index
		}
	}
	return -1
}

// Modifying element
// [2, 4, 6, 8, 10] to 42 and prints the modified array.
func Modifying_E(arr []int) {
	arr[3] = 71
	fmt.Println(arr)
}

// Replacing elements
func replace_E(arr []int, old, new int) []int {
	for index, _ := range arr {
		if arr[index] == old {
			arr[index] = new
		}
	}
	return arr
}

// Edge case
func EdgeCase(arr []int, search int) int {
	for index, _ := range arr {
		if arr[index] == search {
			return search
		}
	}
	return -1

}
