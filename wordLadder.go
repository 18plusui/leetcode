// Source : https://oj.leetcode.com/problems/word-ladder/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Given two words (start and end), and a dictionary, find the length of shortest 
* transformation sequence from start to end, such that:
* 
* Only one letter can be changed at a time
* Each intermediate word must exist in the dictionary
* 
* For example,
* 
* Given:
* start = "hit"
* end = "cog"
* dict = ["hot","dot","dog","lot","log"]
* 
* As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
* return its length 5.
* 
* Note:
* 
* Return 0 if there is no such transformation sequence.
* All words have the same length.
* All words contain only lowercase alphabetic characters.
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
	"strings"
)

var (
	count int
	// tempMap map[string]int
)

func wordLadder(start string, end string, inputArr []string, count int) {

	tempMap := make(map[string]int)
	tempMap[start] = count

	wordArr := strings.Split(start, "")

	for i := 0; i < len(inputArr); i++ {
		if exist(wordArr, 0, end, 0) == 2 {
			fmt.Println("length is : ", tempMap[start]+1)
			fmt.Println("success")
			break
		}

		if exist(wordArr, 0, inputArr[i], 0) == 2 {
			count++
			tempMap[inputArr[i]] = count
			transfer := inputArr[i]
			inputArr[i] = ""
			wordLadder(transfer, end, inputArr, count)
		}

	}

}

func exist(inputArr []string, index int, match string, count int) int {

	for i := 0; i < len(inputArr); i++ {

		if index == len(inputArr) {
			break
		}

		if strings.ContainsAny(inputArr[index], match) {
			count++
		}

		index++

	}
	return count

}

func main() {

	inputArr := []string{"hot", "dot", "dog", "lot", "log"}
	wordLadder("hit", "cog", inputArr, 1)
}
