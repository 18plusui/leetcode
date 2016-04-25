// Source : https://oj.leetcode.com/problems/implement-strstr/
// Author : 18plusui
// Date   : 2016-04-25

/**********************************************************************************
 *
 * Implement strStr().
 *
 * Returns a pointer to the first occurrence of needle in haystack, or null if needle is not part of haystack.
 *
 **********************************************************************************/

package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	index, i, j := 0, 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			index++
			i = index
			j = 0
		}
	}

	if j == len(needle) {
		return index
	}
	return -1
}

func main() {
	fmt.Println(strStr("abcd", "a"))
}
