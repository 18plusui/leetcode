// Source : https://oj.leetcode.com/problems/add-binary/
// Author : 18plusui
// Date   : 2016-04-02

/**********************************************************************************
*
* Given two binary strings, return their sum (also a binary string).
*
* For example,
* a = "11"
* b = "1"
* Return "100".
*
*
**********************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func addBinary(a, b string) {
	var alen, blen int = len(a), len(b)
	var plus bool
	var add string
	var result string

	for alen > 0 || blen > 0 {

		var ha int = threeCul(a, alen)
		var hb int = threeCul(b, blen)
		var hc int
		if plus {
			hc = 1
		} else {
			hc = 0
		}

		if ha+hb+hc == 3 {

			plus = true
			add = "1"

		} else if ha+hb+hc == 2 {

			add = "0"
			plus = true

		} else {
			add = strconv.Itoa(ha + hb + hc)
			plus = false
		}

		result += add

		alen--
		blen--

	}
	if plus {
		result += "1"
	}

	printReslut(result)

}

func threeCul(val string, index int) int {
	if index <= 0 {
		return 0
	}
	retInt, _ := strconv.Atoi(val[index-1 : index])
	return retInt
}

func printReslut(s string) {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))

}

func main() {
	a := "11"
	b := "1"
	addBinary(a, b)
}
