// Source : https://oj.leetcode.com/problems/regular-expression-matching/
// Author : 18plusui
// Date   : 2016-04-09

/**********************************************************************************
*
* Implement regular expression matching with support for '.' and '*'.
*
* '.' Matches any single character.
* '*' Matches zero or more of the preceding element.
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
* isMatch("aa", "a*") → true
* isMatch("aa", ".*") → true
* isMatch("ab", ".*") → true
* isMatch("aab", "c*a*b") → true
*
**********************************************************************************/

package main

import (
	"fmt"
)

func isMatch(s, p string) bool {

	if p == "" {
		return s == ""
	}

	if len(p) == 1 || (len(p) > 1 && string(p[1]) != "*") {
		if s == "" || (string(p[0]) != "." && string(s[0]) != string(p[0])) {
			return false
		}

		return isMatch(s[1:], p[1:])

	} 

	slen := len(s)
	i := -1

	for i < slen && (i < 0 || string(p[0]) == "." || p[0] == s[i]) {
			if len(p) >= 2 {
				if isMatch(s[i+1:], p[2:]) {
					return true
				}
				i++	
			}else{
				if isMatch(s[i+1:], p[1:]) {
					return true
				}
				i++
			}
	}

	return false

}

func main() {
	fmt.Println(isMatch("aa", "."))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))

}
