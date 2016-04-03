// Source : https://leetcode.com/problems/add-digits/
// Author : 18plusui
// Date   : 2016-04-3

/**********************************************************************************
 * Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
 *
 * For example:
 * Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.
 *
 * Follow up:
 * 	Could you do it without any loop/recursion in O(1) runtime?
 *
 **********************************************************************************/
package main

import (
	"fmt"
)

func addDigit(num int) int {
	for num > 9 {
		num = num/10 + num%10
	}
	return num
}

func main() {

	fmt.Print(addDigit(101))

}
