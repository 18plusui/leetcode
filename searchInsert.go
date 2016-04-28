// Source : https://oj.leetcode.com/problems/search-insert-position/
// Author : 18plusui
// Date   : 2016-04-28

/**********************************************************************************
*
* Given a sorted array and a target value, return the index if the target is found.
* If not, return the index where it would be if it were inserted in order.
*
* You may assume no duplicates in the array.
*
* Here are few examples.
* [1,3,5,6], 5 → 2
* [1,3,5,6], 2 → 1
* [1,3,5,6], 7 → 4
* [1,3,5,6], 0 → 0
*
**********************************************************************************/

package main

import (
	"fmt"
)

func binarySearch(nums []int, target int) int {
	n := len(nums)
	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return low
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	return binarySearch(nums, target)
}

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, 2))
}
