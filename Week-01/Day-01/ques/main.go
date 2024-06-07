package main

import "fmt"

func main() {
	// array initialization:
	// arr := [5]int{1, 2, 3, 4, 5} // this is array initliazation
	arr := []int{22, 71, 69, 7, 33, 1} // this is called slice in Golang without specifying size. Dynamic size

	fmt.Println("Array: ", arr)

	// Traversing array:
	traverse(arr)

	// Inserting an Element:
	index := 2
	element := 22
	fmt.Printf("Inserting %d in the index %d : \n", element, index)
	fmt.Println(insArr(arr, index, element))

	// Delete Element:
	idx := 0
	fmt.Printf("Deleting element in the index %d : \n", idx)

	DeleteElement := delArr(arr, idx)
	fmt.Println(DeleteElement)

	// Access Element:
	key := 71
	result := AccessElement(arr, key)

	fmt.Println(arr)

	if result != -1 {
		fmt.Printf("%d is  found in the Array\n", key)
	} else {
		fmt.Println("Key not found in the Array :)")
	}

	fmt.Println(arr)

	// MAX
	fmt.Println("Maximum element in the Array: ")
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
	fmt.Print("Original Array: ", arr)
	reverseArray := reverse(arr)
	fmt.Println("Reverse Array: ", reverseArray)
}

// Array Basics

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

// Sum of Elements:
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


