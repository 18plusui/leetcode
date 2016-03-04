// Source : https://oj.leetcode.com/problems/valid-number/
// Author : 18plusui
// Date   : 2016-03-04

/**********************************************************************************
*
* Validate if a given string is numeric.
*
* Some examples:
* "0" => true
* "   0.1  " => true
* "abc" => false
* "1 a" => false
* "2e10" => true
*
* Note: It is intended for the problem statement to be ambiguous.
*       You should gather all requirements up front before implementing one.
*
**********************************************************************************/

package main

import (
	"fmt"
	"regexp"
)

var numRe = regexp.MustCompile(`^[0-9]`)

func isNumber(s string) bool {
	var dot bool
	var hasE bool
	var curr int
	var slen int = len(s)
	for isSpace(s[curr : curr+1]) {

		curr++
		if curr == slen {
			return false
		}
	}

	if s[curr:curr+1] == "+" || s[curr:curr+1] == "-" {
		curr++
	}

	for s[curr:curr+1] != "0" {
		curr++
		if s[curr:curr+1] == "." {
			if hasE == true || dot == true {
				return false
			}
			if !isDigit(s[curr : curr+1]) {
				return false
			}
			dot = true
			continue

		}

		if s[curr:curr+1] == "e" {
			if hasE == true {
				return false
			}
			curr++
			if s[curr:curr+1] == "+" || s[curr:curr+1] == "-" {
				curr++
			}
			if !isDigit(s[curr : curr+1]) {
				return false
			}
			hasE = true
			continue
		}

		if isSpace(s[curr : curr+1]) {
			for s[curr:curr+1] != "0" {

				if !isSpace(s[curr : curr+1]) {
					return false
				}
				curr++
			}
		}

		if !isDigit(s[curr : curr+1]) {
			return false
		}

	}

	return true
}

func isDigit(c string) bool {
	return numRe.MatchString(c)
}

func isSpace(c string) bool {
	if c == " " {
		return true
	}
	return false
}

func main() {

	fmt.Println(isNumber("0.1"))
	fmt.Println(isNumber("   0.1  "))
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber("abc"))
	fmt.Println(isNumber("1 a"))
	fmt.Println(isNumber("2e10"))

}
