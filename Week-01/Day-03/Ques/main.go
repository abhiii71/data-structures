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

	//  Update 2D Array Element:
	rowIdx = 1
	colIdx = 2
	newValue := 71
	err = updateElement(arr, rowIdx, colIdx, newValue)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Array with Updated Element: ", arr)
	}

	// Traverse 2D Array:
	fmt.Println("Traversed Array: ")
	traverseArr(arr)

	// Sum 2D Array:
	fmt.Printf("Sum of 2D Array: %d \n", sum2Darr(arr))

	// Transpose Matrix:
	fmt.Printf("Transposed Matrix: %d \n", transposeMatrix(arr))

	// Add Matrices:
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrix2 := [][]int{
		{7, 8, 9},
		{10, 11, 12},
	}
	result, err := addMatrix(matrix1, matrix2)
	if err != nil {
		fmt.Printf("Error %s", err)
	} else {
		fmt.Println("Adding two matrix: ", result)
	}

	//
	result, err = subMatrix(matrix1, matrix2)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println("Subtracting two matrices: ", result)
	}

	// Multiply Matrices:
	matrix1 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrix2 = [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	result, err = multiplyMatrices(matrix1, matrix2)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println("Multiplying two matrices: ", result)
	}

	// 2D Array Row Sum:
	arr = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rowSum := rowSumArr(arr)
	fmt.Println("2D Array Row sum: ", rowSum)

	// 2D Array Column Sum:
	colSum := colSumArr(arr)
	fmt.Println("2D Array Column sum: ", colSum)

	// Find Max in 2D Array:
	max := maxElement(arr)
	fmt.Printf("Max Element in the 2D Array: %d \n", max)

	// Find Min in 2D Array:
	min := minElement(arr)
	fmt.Printf("Min Element in the 2D Array: %d \n", min)

	// Diagonal Sum:
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	diagonalSum := diagonalSum(matrix)
	fmt.Printf("Diagonal Sum of the Matrix: %d\n", diagonalSum)

	// Check Symmetry:
	symmetricMatrix := [][]int{
		{1, 2, 3},
		{2, 4, 5},
		{3, 5, 6},
	}

	isSymmetric := checkSymmetry(symmetricMatrix)
	fmt.Printf("Is the matrix symmetric ? %t\n", isSymmetric)
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

// Update 2D Array Element: Update an element in a 2D array.
func updateElement(arr [][]int, rowIdx, colIdx, newValue int) error {
	if rowIdx < 0 || rowIdx >= len(arr) {
		return fmt.Errorf("row index %d out of bounds", rowIdx)
	}

	if colIdx < 0 || rowIdx >= len(arr) {
		return fmt.Errorf("column index %d out of bounds", colIdx)
	}

	arr[rowIdx][colIdx] = newValue
	return nil
}

// Traverse 2D Array: Traverse a 2D array and print its elements.
func traverseArr(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

// Sum 2D Array: Sum all elements in a 2D array.
func sum2Darr(arr [][]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			sum = sum + arr[i][j]
		}
	}
	return sum
}

// Transpose Matrix: Transpose a given matrix.
func transposeMatrix(arr [][]int) [][]int {
	numRow := len(arr)
	numCols := len(arr[0])

	// Create a new matrix with swapped dimensions
	transposed := make([][]int, numCols)
	for i := range transposed {
		transposed[i] = make([]int, numRow)
	}

	// Transpose elements
	for i := 0; i < numRow; i++ {
		for j := 0; j < numCols; j++ {
			transposed[j][i] = arr[i][j]
		}
	}

	return transposed
}

// Add Matrices: Add two matrices.
func addMatrix(matrix1, matrix2 [][]int) ([][]int, error) {
	numRow1, numCol1 := len(matrix1), len(matrix1[0])
	numRow2, numCol2 := len(matrix2), len(matrix2[0])

	// check if dimensions are equal
	if numRow1 != numRow2 || numCol1 != numCol2 {
		return nil, fmt.Errorf("metrix must have the same dimensions")
	}

	// Create a new metrix to store the result
	result := make([][]int, numRow1)
	for i := range result {
		result[i] = make([]int, numCol1)
	}

	// Add corresponding elements
	for i := 0; i < numRow1; i++ {
		for j := 0; j < numCol1; j++ {
			result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	return result, nil
}

// Subtract Matrices: Subtract one matrix from another.
func subMatrix(matrix1, matrix2 [][]int) ([][]int, error) {
	numRow1, numCol1 := len(matrix1), len(matrix1[0])
	numRow2, numCol2 := len(matrix2), len(matrix2[0])

	// Check if the dimensions of both matrices are equal
	if numRow1 != numRow2 || numCol1 != numCol2 {
		return nil, fmt.Errorf("matrices must have the same dimensions")
	}

	// Create a new matrix to store the result
	result := make([][]int, numRow1)
	for i := range result {
		result[i] = make([]int, numCol1)
	}

	// Subtract corresponding elements
	for i := 0; i < numRow1; i++ {
		for j := 0; j < numCol1; j++ {
			result[i][j] = matrix1[i][j] - matrix2[i][j]
		}
	}
	return result, nil
}

// Multiply Matrices: Multiply two matrices.
func multiplyMatrices(matrix1, matrix2 [][]int) ([][]int, error) {
	numRow1, numCol1 := len(matrix1), len(matrix1[0])
	numRow2, numCol2 := len(matrix2), len(matrix2[0])

	if numCol1 != numRow2 {
		return nil, fmt.Errorf("number of columns in matrix1 must be equal to number of rows in matrix2")
	}

	result := make([][]int, numRow1)
	for i := range result {
		result[i] = make([]int, numCol2)
	}

	// Multiply corresponding elements
	for i := 0; i < numRow1; i++ {
		for j := 0; j < numCol2; j++ {
			// Iterate up to numCol1, the number of columns in the first matrix
			for k := 0; k < numCol1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result, nil
}

// 2D Array Row Sum: Calculate the sum of each row in a 2D array.
func rowSumArr(arr [][]int) []int {
	// Initialize a slice to store the sum of each row
	rowSum := make([]int, len(arr))

	// Iterate over each row
	for i, row := range arr {
		// Initialize sum for the current row
		sum := 0

		for _, element := range row {
			sum += element
		}

		// Store the sum of the current row
		rowSum[i] = sum
	}
	// Return the slice containing in the sum of each row
	return rowSum
}

// 2D Array Column Sum: Calculate the sum of each column in a 2D array.
func colSumArr(arr [][]int) []int {
	colSum := make([]int, len(arr))

	numRows := len(arr)
	numCols := len(arr[0])

	for j := 0; j < numCols; j++ {
		sum := 0

		for i := 0; i < numRows; i++ {
			sum += arr[i][j]
		}

		colSum[j] = sum
	}
	return colSum
}

// Find Max in 2D Array: Find the maximum element in a 2D array.
func maxElement(arr [][]int) int {
	max := arr[0][0]

	for _, row := range arr {
		for _, element := range row {
			if element > max {
				max = element
			}
		}
	}
	return max
}

// Find Min in 2D Array: Find the minimum element in a 2D array
func minElement(arr [][]int) int {
	min := arr[0][0]

	for _, row := range arr {
		for _, element := range row {
			if element < min {
				min = element
			}
		}
	}
	return min
}

// Diagonal Sum: Calculate the sum of the diagonal elements of a square matrix.
func diagonalSum(matrix [][]int) int {
	sum := 0

	for i := 0; i < len(matrix); i++ {
		sum += matrix[i][i]
	}
	return sum
}

// Check Symmetry: Check if a given matrix is symmetric.
func checkSymmetry(matrix [][]int) bool {
	transposed := transposeMatrix(matrix)

	numRows := len(matrix)
	numCols := len(matrix[0])
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if matrix[i][j] != transposed[i][j] {
				return false
			}
		}
	}
	return true
}
