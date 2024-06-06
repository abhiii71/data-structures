package main

import (
	"fmt"
	"sort"
)

func main() {
	// Problem 1
	arr := []int{22, 7, 13, 71, 69}
	k := 2
	fmt.Printf("%dth smallest element is %d\n", k, kth(arr, k))

	// Problem 2
	intervals := []Interval{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	mergedIntervals := merge(intervals)
	fmt.Println("Merged intervals: ", mergedIntervals)

	// Probelm 3
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	kk := 3
	fmt.Printf("Maximum sum of subarray of size %d is %d\n", kk, maxSumSubarray(arr1, kk))

	// Problem 4
	nums := []int{1, 2, 3, 7, 5}
	target := 12
	result := subarraySum(nums, target)
	fmt.Println("Subarray with given sum: ", result)
}

// Problem 1  Find Kth smallest/largest elements
func kth(arr []int, k int) int {
	sort.Ints(arr)
	return arr[k-1]
}

// Problem 2 Merging Intervals
type Interval struct {
	Start, End int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}
	for _, interval := range intervals[1:] {
		n := len(merged) - 1
		if merged[n].End >= interval.Start {
			merged[n].End = max(merged[n].End, interval.End)
		} else {
			merged = append(merged, interval)
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Problem 3
func maxSumSubarray(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return -1
	}

	maxSum, windowSum := 0, 0

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum = windowSum

	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// Problem 4
func subarraySum(nums []int, target int) []int {
	prefixSums := map[int]int{0: -1}
	currentSum := 0

	for i, num := range nums {
		currentSum += num
		if j, ok := prefixSums[currentSum-target]; ok {
			return nums[j+1 : i+1]
		}
		prefixSums[currentSum] = i
	}
	return nil
}
