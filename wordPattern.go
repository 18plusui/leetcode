// Source : https://leetcode.com/problems/word-pattern/
// Author : 18plusui
// Date   : 2016-03-07

/*************************************************************************************** 
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 *  Here follow means a full match, such that there is a bijection between a letter in 
 * pattern and a non-empty word in str.
 * 
 * Examples:
 * 
 * pattern = "abba", str = "dog cat cat dog" should return true.
 * pattern = "abba", str = "dog cat cat fish" should return false.
 * pattern = "aaaa", str = "dog cat cat dog" should return false.
 * pattern = "abba", str = "dog dog dog dog" should return false.
 * 
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains lowercase 
 * letters separated by a single space.
 * 
 * Credits:Special thanks to @minglotus6 for adding this problem and creating all test 
 * cases.
 *               
 ***************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	patternMap := make(map[string]string)
	strMap := make(map[string]string)

	plen := len(pattern)
	strArr := strings.Split(str, " ")
	patternArr := strings.Split(pattern, "")

	if plen != len(strArr) {
		return false
	}

	for i := 0; i < plen; i++ {

		if _, ok := patternMap[patternArr[i]]; !ok {

			patternMap[patternArr[i]] = strArr[i]

		}

		if _, ok := strMap[strArr[i]]; !ok {

			strMap[strArr[i]] = patternArr[i]

		}

		if patternMap[patternArr[i]] != strArr[i] || strMap[strArr[i]] != patternArr[i] {
			fmt.Println("pattern map : ", patternMap)
			fmt.Println("string map  : ", strMap)
			return false
		}
	}
	fmt.Println("pattern map : ", patternMap)
	fmt.Println("string map  : ", strMap)
	return true
}

func main() {
	fmt.Println(wordPattern("ABBC", "dog cat cat dog"))
}
