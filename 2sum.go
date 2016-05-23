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

func twoSum(nums []int, target int) []int {
	var result []int

	numslen := len(nums)
	for i := 0; i < numslen; i++ {
		for j := i + 1; j < numslen; j++ {
			if target-nums[i] == nums[j] {

				result = append(result, i)
				result = append(result, j)

				return result
			}

		}
	}
	return result
}
