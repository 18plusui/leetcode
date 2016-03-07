// Source : https://oj.leetcode.com/problems/4sum/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? 
* Find all unique quadruplets in the array which gives the sum of target.
* 
* Note:
* 
* Elements in a quadruplet (a,b,c,d) must be in non-descending order. (ie, a ≤ b ≤ c ≤ d)
* The solution set must not contain duplicate quadruplets.
* 
*     For example, given array S = {1 0 -1 0 -2 2}, and target = 0.
* 
*     A solution set is:
*     (-1,  0, 0, 1)
*     (-2, -1, 1, 2)
*     (-2,  0, 0, 2)
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
	"sort"
	"sync/atomic"
)

var inputArr []int
var arrLen int
var ops uint64 = 0
var sortArr []int
var resNum [4]int

func threeSum(inputArr []int, target int, currInput int) []int {

	if !sort.IntsAreSorted(inputArr) {
		// sort inputArr
		sort.Ints(inputArr)
		// fmt.Println(inputArr)
		sortArr = inputArr
	}

	//get this array length
	arrLen = len(sortArr)

	for i := 0; i < arrLen-2; i++ {
		// fmt.Println("loop time : ", i)
		if i > 0 && sortArr[i-1] == sortArr[i] {
			//if near equal ,skip this time loop
			continue
		}

		currNum := sortArr[i]

		low := i + 1
		high := arrLen - 1

		for low < high {
			lowNum := sortArr[low]
			highNum := sortArr[high]

			sumNum := currNum + lowNum + highNum

			if sumNum == target {
				resNum[0] = currNum
				resNum[1] = lowNum
				resNum[2] = highNum
				resNum[3] = currInput
				printer(resNum, target)
				low++

			} else if sumNum > 0 {

				for high > 0 && sortArr[high] == sortArr[high-1] {
					high--
				}

				high--

			} else {

				for low < arrLen && sortArr[low] == sortArr[low+1] {

					low++
				}
				low++
			}

		}
	}

	return inputArr
}

func fourSum(inputArr []int, target int) {
	if len(inputArr) <= 4 {
		// printer(currNum, lowNum, highNum, target)
	}

	for i := 0; i < len(inputArr)-3; i++ {

		if i > 0 && inputArr[i-1] == inputArr[i] {
			continue
		}

		threeSum(inputArr[i+1:len(inputArr)], target-inputArr[i], inputArr[i])

	}

}

func printer(resNum [4]int, target int) {
	atomic.AddUint64(&ops, 1)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("--------------Begin--------------")
	fmt.Println("Math time : ", opsFinal)
	fmt.Println("result set is : ", resNum[0], " , ", resNum[1], " , ", resNum[2], " , ", resNum[3])
	fmt.Println("---------------End---------------")
	fmt.Println("")
}

func main() {
	inputArr = []int{1, 2, 4, 6, 9, -1, -2, -5, -8, 0}
	fourSum(inputArr, 0)

}
