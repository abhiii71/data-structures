package main

import (
	"fmt"
	"strings"
)

func main() {
	// Daily Temperatures (LeetCode #739)

	fmt.Println("Daily Temperatures (LeetCode #739)")
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))

	// Next Greater Element I (LeetCode #496)

	fmt.Println("Next Greater Element I (LeetCode #496)")
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(nums1, nums2)) // [-1 3 -1]

	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2)) // [3 -1]

	//  Simplify Path (LeetCode #71)
	fmt.Println(" Simplify Path (LeetCode #71)")
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))        // "/"
	fmt.Println(simplifyPath("/home//foo/")) // "/home/foo"
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

/*
Daily Temperatures (LeetCode #739)
Given a list of daily temperatures, you need to determine how many days you would have to wait until a warmer temperature for each day in the input list. If there is no future day with a warmer temperature, return 0 for that day.

Example
For the input list of temperatures:

T = [73, 74, 75, 71, 69, 72, 76, 73]
The output should be:
[1, 1, 4, 2, 1, 1, 0, 0]

Explanation:

For the first temperature 73, you need to wait 1 day for 74.
For the second temperature 74, you need to wait 1 day for 75.
For the third temperature 75, you need to wait 4 days for 76.
And so on.
Solution Explanation
The solution involves using a stack to keep track of the indices of the temperatures. Here's a step-by-step breakdown of the approach:

Initialization:

n is the length of the input list T.
result is a list initialized with zeros, which will store the number of days to wait for a warmer temperature.
stack is used to store the indices of the temperatures.
Iterate Through Temperatures:

For each temperature, check if it is warmer than the temperature corresponding to the index at the top of the stack.
If it is warmer, it means we have found the next warmer day for the day at the index on the top of the stack.
Updating Results:

Pop the index from the stack and update the result list with the difference between the current index and the popped index (which gives the number of days to wait).
Continue this process until the stack is empty or the current temperature is not warmer than the temperature at the index on the top of the stack.
Push Current Index:

Push the current index onto the stack.
This index might be the next day for a future warmer temperature comparison.
Return the Result:

Once the loop completes, the result list contains the required number of days to wait for each day.
*/
func dailyTemperatures(T []int) []int {
	n := len(T)
	result := make([]int, n)
	stack := []int{} // store indices

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
You are given two arrays nums1 and nums2 without duplicates, where nums1 is a subset of nums2. For each element in nums1, you need to find the next greater number in nums2. 
The next greater number for an element x in nums1 is the first number to the right of x in nums2 that is greater than x. If it does not exist, return -1 for that element.

Example
Given the arrays:

nums1 = [4, 1, 2]
nums2 = [1, 3, 4, 2]

The output should be:
[-1, 3, -1]

Explanation:

For 4 in nums1, there is no greater element to the right in nums2.
For 1 in nums1, the next greater element is 3.
For 2 in nums1, there is no greater element to the right in nums2.
Given the arrays:

nums1 = [2, 4]
nums2 = [1, 2, 3, 4]
The output should be:

[3, -1]
Explanation:

For 2 in nums1, the next greater element is 3.
For 4 in nums1, there is no greater element to the right in nums2.
Solution Explanation
The solution involves using a stack and a hash map to efficiently find the next greater elements. Here's the step-by-step approach:

Initialize Data Structures:

result list to store the next greater elements for nums1.
nextGreater hash map to store the next greater element for each number in nums2.
stack to keep track of the elements for which we haven't yet found the next greater element.
Process nums2:

Iterate over each number in nums2.
For each number, if it is greater than the number at the top of the stack, pop the stack and set this number as the next greater element for the popped number in the nextGreater map.
Push the current number onto the stack.
This ensures that we are always comparing the current number with the most recent numbers that haven't yet found their next greater elements.
Populate Result for nums1:

For each number in nums1, check if it has a next greater element in the nextGreater map.
If it does, add that to the result list. Otherwise, add -1.

*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	nextGreater := make(map[int]int)
	stack := []int{}

	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			nextGreater[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for i, num := range nums1 {
		if val, ok := nextGreater[num]; ok {
			result[i] = val
		} else {
			result[i] = -1
		}
	}
	return result
}

/*
Simplify Path (LeetCode #71)
Given an absolute path for a file in Unix-style (i.e., a string path starting with a /), your task is to simplify the path. The goal is to convert the given path to its canonical path.

The canonical path should follow these rules:

The path starts with a single slash /.
Any two directories are separated by a single slash /.
The path does not end with a trailing /.
The path only contains the directories on the path from the root directory to the target file or directory (i.e., no . or ..).
Special components in the path:

. refers to the current directory, so it can be ignored.
.. refers to the parent directory, so it moves up one level in the directory hierarchy.
Any other name is a directory name.
Example
Given the path:

"/home/"
The output should be:

"/home"
Given the path:
"/../"

The output should be:
"/"

Explanation:

.. refers to the parent directory, but since the current directory is the root /, moving up one level still results in /.
Given the path:

"/home//foo/"
The output should be:

"/home/foo"
Solution Explanation
To simplify the Unix-style file path, we can use a stack to handle the directory traversal, especially when encountering .. which requires moving up to the parent directory. Here's the approach step-by-step:

Split the Path:

Split the path by / to get the components (directories, . and ..).
Initialize a Stack:

Use a stack to store the valid directory names.
Process Each Component:

If the component is empty or ., continue to the next component.
If the component is .., pop from the stack if it's not empty (to move up to the parent directory).
Otherwise, push the component onto the stack (as it's a valid directory name).
Construct the Simplified Path:

Join the components in the stack with / to form the canonical path.
Ensure the path starts with /.
*/

func simplifyPath(path string) string {
	components := strings.Split(path, "/")
	stack := []string{}

	for _, component := range components {
		if component == "" || component == "." {
			continue
		}
		if component == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, component)
		}
	}
	return "/" + strings.Join(stack, "/")
}
