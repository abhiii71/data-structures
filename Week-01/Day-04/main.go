package main

import "fmt"

func main() {
	fmt.Println("DAY - 4")

	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
	fmt.Println("Bubble Sort : ", arr)

	arr1 := []int{23, 71, 10, 5, 2}
	insertionSort(arr1)
	fmt.Println("Insertion Sort: ", arr1)

	arr2 := []int{64, 25, 12, 22, 11}
	selectionSort(arr2)
	fmt.Println("Selection Sort: ", arr2)

	arr3 := []int{10, 50, 30, 70, 80, 60, 20, 90, 40}
	target := 30
	result := lSearch(arr3, target)
	if result != -1 {
		fmt.Println("Linear Search: ", "Element Found At Index: ", result)
	} else {
		fmt.Println("Linear Search: ;) Element not found in Array.")
	}
}

// Bubble Sort
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1 // 0 -  1 - 2

		for j >= 0 && arr[j] > key { // false for  0 index  // true for 1st index
			arr[j+1] = arr[j] // arr[2] = arr[1]
			j = j - 1         // j = 0 ,1
		}
		arr[j+1] = key //
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func lSearch(arr []int, target int) int {
	for key, value := range arr {
		if value == target {
			return key
		}
	}
	return -1
}
