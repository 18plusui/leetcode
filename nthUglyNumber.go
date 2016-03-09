// Source : https://leetcode.com/problems/ugly-number-ii/
// Author : 18plusui
// Date   : 2016-03-09

/***************************************************************************************
 *
 * Write a program to find the n-th ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For
 * example, 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
 *
 * Note that 1 is typically treated as an ugly number.
 *
 *   The naive approach is to call isUgly for every number until you reach the nth one.
 * Most numbers are not ugly. Try to focus your effort on generating only the ugly ones.
 *
 *   An ugly number must be multiplied by either 2, 3, or 5 from a smaller ugly number.
 *
 *   The key is how to maintain the order of the ugly numbers. Try a similar approach
 * of merging from three sorted lists: L1, L2, and L3.
 *
 *   Assume you have Uk, the kth ugly number. Then Uk+1 must be Min(L1 * 2, L2 * 3, L3
 * * 5).
 *
 ***************************************************************************************/

package main

import (
	"fmt"
)

func isNthUgly(n int) []int {

	var x, y, z int
	var res = make([]int, 0, 10)
	var culVal int
	//init slice res first value = 1
	res = append(res, 1)

	for len(res) < n {
		culVal = min(res[x]*2, res[y]*3, res[z]*5)
		fmt.Println("cul value : ", culVal)
		if culVal == res[x]*2 {
			x++
		}
		if culVal == res[y]*3 {
			y++
		}
		if culVal == res[z]*5 {
			z++
		}
		res = append(res, culVal)
	}

	return res
}

func min(a, b, c int) int {
	return min2(min2(a, b), c)
}

func min2(a, b int) int {

	if a < b {
		return a
	}

	return b
}

func main() {

	fmt.Println(isNthUgly(10))
}
