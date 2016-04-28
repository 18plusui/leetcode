// Source : https://oj.leetcode.com/problems/search-for-a-range/
// Author : 18plusui
// Date   : 2016-04-28

/**********************************************************************************
*
* Given a sorted array of integers, find the starting and ending position of a given target value.
*
* Your algorithm's runtime complexity must be in the order of O(log n).
*
* If the target is not found in the array, return [-1, -1].
*
* For example,
* Given [5, 7, 7, 8, 8, 10] and target value 8,
* return [3, 4].
*
**********************************************************************************/

package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	n := len(nums)
	pos := binarySearch(nums, 0, n-1, target)

	var result []int
	low, high := -1, -1
	if pos >= 0 {
		high = pos
		low = high

		l := binarySearch(nums, 0, low-1, target)
		for l >= 0 {
			low = l
			l = binarySearch(nums, 0, low-1, target)
		}

		h := binarySearch(nums, high+1, n-1, target)
		for h >= 0 {
			high = h
			h = binarySearch(nums, high+1, n-1, target)
		}
	}
	result = append(result, low)
	result = append(result, high)

	return result
}

func binarySearch(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
		}

		if target < nums[mid] {
			high = mid - 1
		}
	}
	return -1

}

func main() {
	fmt.Println(searchRange([]int{1, 2, 2}, 2))
}
