package main

import "fmt"

func main() {
	// declaring a 2D array
	var matrix [3][3]int

	// Initialize the 2D array

	matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("2D Array: ", matrix)

	// Traversing elements
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}

	// Matrix Addition
	MatAdd()

	// Matrix Transpose
	MatTranspose()

	// Questions:
	ques()

	// ques 2
	mat1 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mat2 := [3][3]int{{4, 5, 6}, {7, 8, 9}, {1, 2, 3}}
	result := ques2(mat1, mat2)
	fmt.Println("Matrix Addition Result")
	for _, row := range result {
		fmt.Println(row)
	}
}

func MatAdd() {
	// Matrix addition
	mat2 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	mat3 := [2][2]int{
		{5, 6},
		{7, 8},
	}

	var result [2][2]int
	for i := 0; i < len(mat2); i++ {
		for j := 0; j < len(mat3); j++ {
			result[i][j] = mat2[i][j] + mat3[i][j]
		}
	}

	fmt.Println("Matrix Addition Result: ", result)
}

// Matrix Transpose
func MatTranspose() {
	mat := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	var transpose [3][2]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			transpose[j][i] = mat[i][j]
		}
	}
	fmt.Println("Transpose of Matrix: ", transpose)
}

// Questions

func ques() {
	var matrix [4][4]int
	count := 1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = count
			count++
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}
}

func ques2(mat1, mat2 [3][3]int) [3][3]int {
	var result [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = mat1[i][j] + mat2[i][j]
		}
	}
	return result
}

