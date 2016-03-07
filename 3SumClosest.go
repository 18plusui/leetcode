// Source : https://oj.leetcode.com/problems/3sum-closest/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Given an array S of n integers, find three integers in S such that the sum is 
* closest to a given number, target. Return the sum of the three integers. 
* You may assume that each input would have exactly one solution.
* 
*     For example, given array S = {-1 2 1 -4}, and target = 1.
* 
*     The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
	"sort"
	"sync/atomic"
)

var maxInt int = 2147483647
var arrLen int
var ops uint64 = 0
var sortArr []int

func cal(inputArr []int, target int) {

	//make sure are sorted
	if !sort.IntsAreSorted(inputArr) {
		// sort inputArr
		sort.Ints(inputArr)
		fmt.Println(inputArr)
		sortArr = inputArr
	}

	//get this array length
	arrLen = len(sortArr)

	for i := 0; i < arrLen-2; i++ {
		fmt.Println("loop time : ", i)
		if i > 0 && sortArr[i-1] == sortArr[i] {
			fmt.Println("I'm come in")
			continue
		}

		currNum := sortArr[i]

		low := i + 1
		high := arrLen - 1

		for low < high {
			lowNum := sortArr[low]
			highNum := sortArr[high]
			sumNum := currNum + lowNum + highNum

			if sumNum-target == 0 {
				printer(currNum, lowNum, highNum, target)
				low++

			} else {

				if abs(sumNum) < maxInt {

					maxInt = abs(sumNum)
					printer(currNum, lowNum, highNum, sumNum)

				}

				if sumNum-target > 0 {

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
	}
}

func printer(currNum int, lowNum int, highNum int, target int) {
	atomic.AddUint64(&ops, 1)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("--------------Begin--------------")
	fmt.Println("counter : ", opsFinal)
	fmt.Println("first Number : ", currNum, " , ", lowNum, " , ", highNum)
	fmt.Println("target : ", target)
	fmt.Println("---------------End---------------")
	fmt.Println("")
}

func abs(sn int) int {
	if sn > 0 {
		return sn
	}

	return -sn
}

func main() {
	cal([]int{-4, -2, 0, -8, -11, 3, 6, 7}, 0)
	fmt.Println("the last Max Int number is : ", maxInt)
}
