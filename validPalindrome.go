// Source : https://oj.leetcode.com/problems/valid-palindrome/
// Author : 18plusui
// Date   : 2016-03-03

/**********************************************************************************
*
* Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
*
* For example,
* "A man, a plan, a canal: Panama" is a palindrome.
* "race a car" is not a palindrome.
*
* Note:
* Have you consider that the string might be empty? This is a good question to ask during an interview.
*
* For the purpose of this problem, we define empty string as valid palindrome.
*
*
**********************************************************************************/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = removeNoise(s)
	var slen int = len(s)
	for i := 0; i < slen/2; i++ {
		fmt.Println("pre : ", s[i:i+1], "end : ", s[slen-i-1:slen-i])
		if s[i:i+1] != s[slen-i-1:slen-i] {
			return false
		}
	}
	return true

}

func removeNoise(inputStr string) string {

	var validStr = regexp.MustCompile(`^[a-z]`)
	var unnoiseStr string
	var cnt int = 0

	inputStr = strings.ToLower(inputStr)
	for i := 0; i < len(inputStr); i++ {

		if validStr.MatchString(inputStr[i : i+1]) {

			unnoiseStr += inputStr[i : i+1]
			cnt++
		}
	}

	return unnoiseStr

}

func main() {

	inputStr := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(inputStr))

}
