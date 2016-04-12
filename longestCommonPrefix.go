// Source : https://oj.leetcode.com/problems/longest-common-prefix/
// Author : 18plusui
// Date   : 2016-04-13

/********************************************************************************** 
* 
* Write a function to find the longest common prefix string amongst an array of strings.
*                
**********************************************************************************/

package main

import (
	"fmt"
)

func longestPrefix (strs []string) string {
	var  word string
	if len(strs) <= 0 { return word }
    for i := 1; i <= len(strs[0]); i++ {
        w := (strs[0])[:i]
        match := true
        for j := 1; j < len(strs); j++ {
        	if i > len(strs[j]) || w != (strs[j])[:i] {
        		match = false
        		break
        	}
        }
        if !match {
        	return word
        }
        word = w 

    }
    return word
}

func main() {
	s := []string{"abab"}
	fmt.Println(longestPrefix(s))
}
