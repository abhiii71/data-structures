package main

import "fmt"

func main() {
	// Initialize 2D Array: Initialize a 2D array with values.
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(arr)

	// Access 2D Array Element:
	rowIdx := 1
	colIdx := 2
	element, err := accessElement(arr, rowIdx, colIdx)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Printf("Element at arr[%d][%d] is %d\n", rowIdx, colIdx, element)
	}
}

// Access 2D Array Element: Access an element in a 2D array.
func accessElement(arr [][]int, rowIdx, colIdx int) (int, error) {
	if rowIdx < 0 || rowIdx >= len(arr) {
		return 0, fmt.Errorf("row index %d is out  of bounds", rowIdx)
	}
	if colIdx < 0 || colIdx >= len(arr[rowIdx]) {
		return 0, fmt.Errorf("column index %d out of bounds", colIdx)
	}
	return arr[rowIdx][colIdx], nil
}
