// Source : https://oj.leetcode.com/problems/reverse-integer/
// Author : 18plusui
// Date   : 2016-04-08

/**********************************************************************************
*
* Reverse digits of an integer.
*
* Example1: x =  123, return  321
* Example2: x = -123, return -321
*
*
* Have you thought about this?
*
* Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!
*
* > If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.
*
* > Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer,
*   then the reverse of 1000000003 overflows. How should you handle such cases?
*
* > Throw an exception? Good, but what if throwing an exception is not an option?
*   You would then have to re-design the function (ie, add an extra parameter).
*
*
**********************************************************************************/

package main

import (
	"fmt"
)

const (
	INT_MAX int = 2147483647
	INT_MIN int = -2147483648
)

func reverse(x int) int {
	y := 0
	var n int
	for x != 0 {
		n = x % 10
		if y > INT_MAX/10 || y < INT_MIN/10 {
			return 0
		}
		y = y*10 + n
		x /= 10
	}
	fmt.Println(y)
	return y
}

func main() {
	reverse(123155123)
}
