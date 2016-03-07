// Source : https://oj.leetcode.com/problems/3sum/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? 
* Find all unique triplets in the array which gives the sum of zero.
* 
* Note:
* 
* Elements in a triplet (a,b,c) must be in non-descending order. (ie, a ≤ b ≤ c)
* The solution set must not contain duplicate triplets.
* 
*     For example, given array S = {-1 0 1 2 -1 -4},
* 
*     A solution set is:
*     (-1, 0, 1)
*     (-1, -1, 2)
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
	"sort"
	"sync/atomic"
)

var arrLen int
var ops uint64 = 0
var sortArr []int

func getArr() []int {
	inputArr := []int{-7, -8, -1, -2, -4, 1, 4, 9, 3, 6}
	return inputArr
}

func cal() {
	// call get array func
	inputArr := getArr()

	//make sure are sorted
	fmt.Println(inputArr)
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

			if sumNum == 0 {
				//2d array use
				atomic.AddUint64(&ops, 1)
				opsFinal := atomic.LoadUint64(&ops)

				//print result
				fmt.Println("--------------Begin--------------")
				fmt.Println("counter : ", opsFinal)
				fmt.Println("first Number : ", currNum)
				fmt.Println("second Number : ", lowNum)
				fmt.Println("third Number : ", highNum)
				fmt.Println("---------------End---------------")
				fmt.Println("")

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

}

func main() {

	cal()

}
