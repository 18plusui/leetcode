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
    var cnt [255]int
    var start int = -1
    var maxLen int = 0
    var length int = len(s)
    for i := 0; i < 255; i++ {
        cnt[i] = -1
    }
    
    for end:= 0; end < length;  {
        if cnt[int(s[end])] == -1 {
            cnt[int(s[end])] = end
            if end - start > maxLen {
                maxLen = end - start
            }
            end++
        } else {
            for i := start + 1; i < cnt[int(s[end])]; i++ {
                cnt[int(s[i])] = -1
            }
            start = cnt[int(s[end])]
            cnt[int(s[end])] = -1
        }
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
