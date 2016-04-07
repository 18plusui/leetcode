// Source : https://oj.leetcode.com/problems/longest-substring-without-repeating-characters/
// Author : 18plusui
// Date   : 2016-04-07

/**********************************************************************************
*
* Given a string, find the length of the longest substring without repeating characters.
* For example, the longest substring without repeating letters for "abcabcbb" is "abc",
* which the length is 3. For "bbbbb" the longest substring is "b", with the length of 1.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	m := make(map[string]int)
	var maxLen int
	var lastRepeatPos int = 0

	for i := 0; i < len(s); i++ {

		if _, ok := m[string(s[i])]; ok && lastRepeatPos < m[string(s[i])] {
			fmt.Println("lastPos : ", m[string(s[i])])
			lastRepeatPos = m[string(s[i])]
			continue
		}

		if _, ok := m[string(s[i])]; !ok {
			m[string(s[i])] = i
		}

		fmt.Println("map : ", m)

	}

	if lastRepeatPos == 0 {
		maxLen = 1
	} else {
		maxLen = lastRepeatPos + 1
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("abcdabcdbb"))
}
