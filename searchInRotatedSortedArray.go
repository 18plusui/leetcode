// Source : https://oj.leetcode.com/problems/search-in-rotated-sorted-array/
// Author : 18plusui
// Date   : 2016-04-28

/**********************************************************************************
*
* Suppose a sorted array is rotated at some pivot unknown to you beforehand.
*
* (i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
*
* You are given a target value to search. If found in the array return its index, otherwise return -1.
*
* You may assume no duplicate exists in the array.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func search(A []int, key int) int {
	n := len(A)
	if n <= 0 {
		return -1
	}

	if n == 1 {
		if A[0] == key {
			return 0
		}
		return -1
	}

	low, high := 0, n-1

	for low <= high {
		if A[low] <= A[high] && (key < A[low] || key > A[high]) {
			return -1
		}

		mid := low + (high-low)/2

		if A[mid] == key {
			return mid
		}

		if A[low] < A[mid] && key >= A[low] && key < A[mid] {
			high = mid - 1
			continue
		}

		if A[mid] < A[high] && key > A[mid] && key <= A[high] {
			low = mid + 1
			continue
		}

		if A[low] > A[mid] {
			high = mid - 1
			continue
		}

		if A[mid] > A[high] {
			low = mid + 1
			continue
		}
	}

	return -1

}

func main() {
	fmt.Println(search([]int{1, 3}, 2))
}
