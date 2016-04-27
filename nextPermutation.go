// Source : https://oj.leetcode.com/problems/next-permutation/
// Author : 18plusui
// Date   : 2016-04-27

/**********************************************************************************
*
* Implement next permutation, which rearranges numbers into the lexicographically next
* greater permutation of numbers.
*
* If such arrangement is not possible, it must rearrange it as the lowest possible order
* (ie, sorted in ascending order).
*
* The replacement must be in-place, do not allocate extra memory.
*
* Here are some examples. Inputs are in the left-hand column and its corresponding outputs
* are in the right-hand column.
*
*   1,2,3 → 1,3,2
*   3,2,1 → 1,2,3
*   1,1,5 → 1,5,1
*
**********************************************************************************/

package main

import (
	"fmt"
)

func nextPermutation(num []int) {
	if len(num) <= 1 {
		return
	}

	for i := len(num) - 1; i > 0; i-- {
		if num[i-1] < num[i] {
			j := len(num) - 1
			for num[i-1] >= num[j] {
				j--
			}

			num[j], num[i-1] = num[i-1], num[j]
			reverse(num[i:])
			fmt.Println(num)
			return
		}

		if i == 1 {
			return
		}
	}

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {

	nextPermutation([]int{1, 2, 3})

}
