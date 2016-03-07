// Source : https://oj.leetcode.com/problems/valid-number/
// Author : 18plusui
// Date   : 2016-03-07

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
	"unicode"
)

func isNumber(s string) bool {

	var dot bool
	var hasE bool
	var curr int
	var slen int = len(s)

	for _, c := range s {

		if unicode.IsSpace(c) {
			curr++
			continue
		}

		break

	}

	if curr == slen {
		return false
	}

	if s[curr] == '+' || s[curr] == '-' {
		curr++
	}

	if !unicode.IsDigit(rune(s[curr])) {
		return false
	} else {
		curr++
	}

	tmpSlice := s[curr:]

	for i := 0; i != len(tmpSlice); i++ {

		if tmpSlice[i] == '.' {
			if hasE == true || dot == true {
				return false
			}

			i++

			if !unicode.IsDigit(rune(tmpSlice[i])) {
				return false
			}

			dot = true
			continue

		}

		if tmpSlice[i] == 'e' {

			if hasE == true {
				return false
			}
			i++
			if tmpSlice[i] == '+' || tmpSlice[i] == '-' {
				i++
			}
			if !unicode.IsDigit(rune(tmpSlice[i])) {
				return false
			}
			hasE = true
			continue
		}

		if unicode.IsSpace(rune(tmpSlice[i])) {

			for _, c := range tmpSlice[i:] {
				if !unicode.IsSpace(c) {
					return false
				}

			}
			return true
		}

		if !unicode.IsDigit(rune(tmpSlice[i])) {
			return false
		}

	}

	return true
}

func main() {

	fmt.Println(isNumber("0.1"))
	fmt.Println(isNumber("   0.1  "))
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber("abc"))
	fmt.Println(isNumber("1 a"))
	fmt.Println(isNumber("2.2e10"))
	fmt.Println(isNumber("-2.2e-4"))

}
