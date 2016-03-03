// Source : https://oj.leetcode.com/problems/valid-parentheses/
// Author : 18plusui
// Date   : 2016-03-03

/**********************************************************************************
*
* Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
* determine if the input string is valid.
*
* The brackets must close in the correct order, "()" and "()[]{}" are all valid
* but "(]" and "([)]" are not.
*
*
**********************************************************************************/

package main

import (
	"fmt"
)

func isValid(inputArr []string) bool {

	arrLen := len(inputArr)
	if arrLen == 0 {
		return false
	}

	var tmpCopy = make([]string, arrLen)
	for i, v := range inputArr {
		if v == "(" || v == "[" || v == "{" {
			tmpCopy[i] = v
		} else if v == ")" || v == "]" || v == "}" {
			if i == 0 || tmpCopy[i-1] == "" {
				return false
			}

			if (v == ")" && tmpCopy[i-1] == "(") || (v == "]" && tmpCopy[i-1] == "[") ||
				(v == "}" && tmpCopy[i-1] == "{") {
				tmpCopy[i] = v
			} else {
				return false
			}
		} else {
			return false
		}
		fmt.Println(tmpCopy)
	}

	return true

}

func main() {
	inputArr := []string{"(", ")", "[", "}", "{", "]"}
	fmt.Println(isValid(inputArr))
}
