package main

import "fmt"

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperature(T))
}

/*
Daily Temperatures (LeetCode #739)

Given a list of daily temperatures, return a list that indicates how many days to wait for a warmer temperature.
Example:
plaintext
Copy code
Input: [73, 74, 75, 71, 69, 72, 76, 73]
Output: [1, 1, 4, 2, 1, 1, 0, 0]
*/
func dailyTemperature(T []int) []int {
	n := len(T)
	result := make([]int, n)

	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return result
}

/*
Next Greater Element I (LeetCode #496)

Given two arrays where one is a subset of the other, find the next greater number for elements of the subset.
Example:
plaintext
Copy code
Input: nums1 = [4, 1, 2], nums2 = [1, 3, 4, 2]
Output: [-1, 3, -1]
*/
func NextGreater(nums1 []int, nums2 []int) []int {
	// 	nextGreater := make(map[int]int)
	//	stack := []int{}

	return nums2
}
