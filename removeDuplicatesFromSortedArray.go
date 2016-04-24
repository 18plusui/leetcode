// Source : https://oj.leetcode.com/problems/remove-duplicates-from-sorted-array/
// Author : 18plusui
// Date   : 2016-04-25

/**********************************************************************************
*
* Given a sorted array, remove the duplicates in place such that each element appear
* only once and return the new length.
*
* Do not allocate extra space for another array, you must do this in place with constant memory.
*
* For example,
* Given input array A = [1,1,2],
*
* Your function should return length = 2, and A is now [1,2].
*
**********************************************************************************/

package main

import (
	"fmt"
)

func removeDuplicates(input []int) int {

	length := len(input)

	if length <= 1 {
		return length
	}

	pos := 0

	for i := 0; i < length-1; i++ {

		if input[i] != input[i+1] {
			pos++
			input[pos] = input[i+1]
		}

	}

	fmt.Println(input[:pos+1])

	return pos + 1

}

func main() {
	input := []int{1, 1, 2, 3, 3, 3, 4, 5}
	length := removeDuplicates(input)
	fmt.Println(length)
}
