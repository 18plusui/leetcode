// Source : https://oj.leetcode.com/problems/longest-valid-parentheses/
// Author : Hao Chen
// Date   : 2014-07-18

/**********************************************************************************
*
* Given a string containing just the characters '(' and ')',
* find the length of the longest valid (well-formed) parentheses substring.
*
* For "(()", the longest valid parentheses substring is "()", which has length = 2.
*
* Another example is ")()())", where the longest valid parentheses substring is "()()",
* which has length = 4.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	var maxlen int
	lastError := -1
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				length := 0
				if len(stack) == 0 {
					length = i - lastError
				} else {
					length = i - stack[len(stack)-1]
				}
				if length > maxlen {
					maxlen = length
				}
			} else {
				lastError = i
			}
		}
	}
	return maxlen
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
