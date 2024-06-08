package main

import (
	"fmt"
)

func main() {
	// array initialization:
	// arr := [5]int{1, 2, 3, 4, 5} // this is array initliazation
	arr := []int{22, 71, 69, 7, 33, 1} // this is called slice in Golang without specifying size. Dynamic size

	fmt.Println("Array: ", arr)

	// Array Length:
	fmt.Println("Length of the Array: ", lenArr(arr))

	// Traversing array:
	traverse(arr)

	// Inserting an Element:
	index := 2
	element := 22
	fmt.Printf("Inserting %d in the index %d : ", element, index)
	fmt.Println(insArr(arr, index, element))

	// Delete Element:
	idx := 0
	fmt.Printf("Deleting element in the index %d : ", idx)

	DeleteElement := delArr(arr, idx)
	fmt.Println(DeleteElement)

	// Update Element:

	indx := 1
	new := 10
	fmt.Print("Updated array: ")
	fmt.Println(updateArr(arr, indx, new))

	// Access Element:
	key := 71
	result := AccessElement(arr, key)

	if result != -1 {
		fmt.Printf("%d is  found in the Array\n", key)
	} else {
		fmt.Println("Key not found in the Array :)")
	}

	// MAX
	fmt.Print("Maximum element in the Array: ")
	fmt.Println(max(arr))

	// MIN
	fmt.Printf("Minimum element in the Array: ")
	minElement := min(arr)
	fmt.Println(minElement)

	// SUM
	fmt.Print("Sum of the Array: ")
	fmt.Println(sum(arr))

	// AVERAGE
	fmt.Print("Average of the Array: ")
	fmt.Println(avg(arr))

	// Reverse Array
	fmt.Print("Original Array: ", arr, "\n")
	reverseArray := reverse(arr)
	fmt.Println("Reverse Array: ", reverseArray)

	// Copy Array
	src := []int{1, 2, 3, 4, 5, 6}
	des := make([]int, lenArr(src))
	copyArr(src, des)
	fmt.Println("source Array: ", src)
	fmt.Println("Destinition Array: ", des)

	// Even Numbers:
	fmt.Println("Even Elements in the Array: ")
	printEven(arr)

	// Odd Numbers:
	fmt.Println("Odd Elements in the Array: ")
	printOdd(arr)

	// Array Multiplication:
	array := []int{3, 6, 9}
	fmt.Print("Multiply each element of an array by 2: ")
	MultiplyArr(array)
	fmt.Println(array)
}

// Array Basics

// Array Length:
func lenArr(arr []int) int {
	count := 0
	for range arr {
		count++
	}
	return count
}

// Traversing array
func traverse(arr []int) {
	fmt.Println("Traversing the Array: ")

	for _, element := range arr {
		fmt.Println(element)
	}
}

// Insert Element:
func insArr(arr []int, index, element int) []int {
	return append(arr[:index], append([]int{element}, arr[index:]...)...)
}

// Delete Element:
func delArr(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// Update Element:
func updateArr(arr []int, indx, new int) []int {
	for index := range arr {
		if index == indx {
			arr[index] = new
		}
	}
	return arr
}

// Access Element:
func AccessElement(arr []int, key int) int {
	for index, element := range arr {
		if element == key {
			return index
		}
	}
	return -1
}

// Find Maximum:
func max(arr []int) int {
	max := arr[0]
	for _, element := range arr {
		if element > max {
			max = element
		}
	}

	return max
}

// Find Minimum:
func min(arr []int) int {
	min := arr[0]
	for _, element := range arr {
		if element < min {
			min = element
		}
	}
	return min
}

// Sum Elements: Sum all elements in an array.
func sum(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	return sum
}

// Average of Elements:
func avg(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	avg := sum / len(arr)
	return avg
}

// Reverse Array:
/*
func reverse(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		left++
		right--
	}
	return arr
}
*/

// Reverse Array:
func reverse(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

// Copy Array: Copy all elements of one array to another
func copyArr(src []int, des []int) {
	if lenArr(src) != lenArr(des) {
		fmt.Println("Length of Source Array doesn't match with Destination Array.")
		return
	}

	for index, element := range src {
		des[index] = element
	}
}

// Even Numbers: Write a function to print all even numbers in an array
func printEven(arr []int) {
	for _, element := range arr {
		if element%2 == 0 {
			fmt.Println(element)
		}
	}
}

// Odd Numbers: Write a function to print all odd numbers in an array
func printOdd(arr []int) {
	for _, element := range arr {
		if element%2 != 0 {
			fmt.Println(element)
		}
	}
}

// Array Multiplication: Multiply each element of an array by 2
func MultiplyArr(arr []int) {
	for index := range arr {
		arr[index] = arr[index] * 2
	}
}
