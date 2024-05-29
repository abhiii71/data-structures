package main

import "fmt"

func main() {
	arr()
	fmt.Println("Now questions:")
	ques()

	// Accessing Element
	ary := []int{10, 20, 30, 40, 50}
	fmt.Println(getElement(ary, 2))

	// Insertion
	fmt.Println(insertElement(ary, 2, 71))

	// deletion
	fmt.Println(deleteElement(ary, 2))

	// traversal
	printElemets(ary)

}

func arr() {
	//array
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array: ", arr)

	//accessing element from the array
	fmt.Println(arr[0])

	// insertion of array
	arr = append(arr[:2], append([]int{25}, arr[2:]...)...)
	fmt.Println("Insertion in the array: ", arr)

	//deletion of the array delete 25 now
	arr = append(arr[:2], arr[3:]...)
	fmt.Println("After Deletion: ", arr)
}

func ques() {
	// Insertion array  at index 3
	arr := []int{1, 2, 3, 4, 5, 6}
	arr = append(arr[:3], append([]int{71}, arr[3:]...)...)
	fmt.Println("Array Insertion: ", arr)

	arr = append(arr[:len(arr)], append([]int{69})...)
	fmt.Println("Array: ", arr)

	// delete the element at index 1 // means ar[1] = 2
	ar := []int{1, 2, 3, 4, 5, 6}
	ar = append(ar[:1], append(ar[2:])...)
	fmt.Println("ar: ", ar)
}

// Accessing Elements
func getElement(arr []int, index int) int {
	if index < 0 || index >= len(arr) {
		return -1
	}
	return arr[index]
}

// Insertion
func insertElement(arr []int, index int, element int) []int {
	return append(arr[:index], append([]int{element}, arr[index:]...)...)
}

// deletion
func deleteElement(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// Array Traversal
func printElemets(arr []int) {
	for _, element := range arr {
		fmt.Println(element)
	}
}
