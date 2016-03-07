// Source : https://oj.leetcode.com/problems/word-break/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Given a string s and a dictionary of words dict, determine if s can be segmented 
* into a space-separated sequence of one or more dictionary words.
* 
* For example, given
* s = "leetcode",
* dict = ["leet", "code"].
* 
* Return true because "leetcode" can be segmented as "leet code".
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
)

func wordBreak(inputStr string, dict []string) bool {

	start := 0
	success := 0

	for i := 0; i <= len(inputStr); i++ {
		tmpSubstr := inputStr[start:i]
		for j := 0; j < len(dict); j++ {
			if tmpSubstr == dict[j] {
				start = i
				success++
				if success == len(dict) {
					return true
				}
			}

		}

	}

	return false
}

func main() {
	fmt.Println(wordBreak("helloworld", []string{"hello", "world"}))

}
