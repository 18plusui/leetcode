// Source : https://oj.leetcode.com/problems/wildcard-matching/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Implement wildcard pattern matching with support for '?' and '*'.
* 
* '?' Matches any single character.
* '*' Matches any sequence of characters (including the empty sequence).
* 
* The matching should cover the entire input string (not partial).
* 
* The function prototype should be:
* bool isMatch(const char *s, const char *p)
* 
* Some examples:
* isMatch("aa","a") → false
* isMatch("aa","aa") → true
* isMatch("aaa","aa") → false
* isMatch("aa", "*") → true
* isMatch("aa", "a*") → true
* isMatch("ab", "?*") → true
* isMatch("aab", "c*a*b") → false
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
)

func isMatch(inputStr string, matchStr string) bool {
	inputStart := 0
	matchStart := 0

	for inputStart != len(inputStr) {
		if matchStr[matchStart:matchStart+1] == "*" {
			matchStart++
			if matchStart == len(matchStr) {
				return true
			}

		} else if matchStr[matchStart:matchStart+1] == "?" ||
			inputStr[inputStart:inputStart+1] == matchStr[matchStart:matchStart+1] {
			matchStart++
			inputStart++
			if matchStart == len(matchStr) && matchStart == len(inputStr) {
				return true

			}
		} else if inputStart != 0 {
			inputStart++

		}

	}

	return false

}

func main() {

	fmt.Println(isMatch("hello", "h??o"))
}
