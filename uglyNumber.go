// Source : https://leetcode.com/problems/ugly-number/
// Author : 18plusui
// Date   : 2016-03-09

/***************************************************************************************
 *
 * Write a program to check whether a given number is an ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For
 * example, 6, 8 are ugly while 14 is not ugly since it includes another prime factor 7.
 *
 * Note that 1 is typically treated as an ugly number.
 *
 ***************************************************************************************/

package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	if num == 1 {
		return true
	}

	for num%2 == 0 {
		num /= 2
		fmt.Println("number 2: ", num)
	}

	for num%3 == 0 {
		num /= 3
		fmt.Println("number 3: ", num)
	}

	for num%5 == 0 {
		num /= 5
		fmt.Println("number 5: ", num)
	}

	return num == 1

}

func main() {
	fmt.Println(isUgly(100))
}
