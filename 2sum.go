// Source : https://oj.leetcode.com/problems/two-sum/
// Author : 18plusui
// Date   : 2016-03-10

/**********************************************************************************
*
* Given an array of integers, find two numbers such that they add up to a specific target number.
*
* The function twoSum should return indices of the two numbers such that they add up to the target,
* where index1 must be less than index2. Please note that your returned answers (both index1 and index2)
* are not zero-based.
*
* You may assume that each input would have exactly one solution.
*
* Input: numbers={2, 7, 11, 15}, target=9
* Output: index1=1, index2=2
*
**********************************************************************************/

package leetcode

// import (
// 	"fmt"
// )

func Sum2ptr(arr *[]int, target int) {

	arrLen := len(*arr)

	for i := 0; i < arrLen; i++ {

		for j := i + 1; j < arrLen; j++ {
			if target-(*arr)[i] == (*arr)[j] {
				// fmt.Println((*arr)[i], (*arr)[j])
			}

		}
	}

}

func Sum2arr(arr []int, target int) {

	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {

		for j := i + 1; j < arrLen; j++ {
			if target-arr[i] == arr[j] {
				// fmt.Println(arr[i], arr[j])
			}

		}
	}

}
