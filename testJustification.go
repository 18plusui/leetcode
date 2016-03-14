// Source : https://oj.leetcode.com/problems/text-justification/
// Author : 18plusui
// Date   : 2016-03-14

/**********************************************************************************
*
* Given an array of words and a length L, format the text such that each line has
* exactly L characters and is fully (left and right) justified.
*
*
* You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
* Pad extra spaces ' ' when necessary so that each line has exactly L characters.
*
* Extra spaces between words should be distributed as evenly as possible.
* If the number of spaces on a line do not divide evenly between words,
* the empty slots on the left will be assigned more spaces than the slots on the right.
*
* For the last line of text, it should be left justified and no extra space is inserted between words.
*
* For example,
* words: ["This", "is", "an", "example", "of", "text", "justification."]
* L: 16.
*
* Return the formatted lines as:
*
* [
*    "This    is    an",
*    "example  of text",
*    "justification.  "
* ]
*
* Note: Each word is guaranteed not to exceed L in length.
*
*
* Corner Cases:
*
* A line other than the last line might contain only one word. What should you do in this case?
* In this case, that line should be left-justified.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func fullJustify(input []string, long int) []string {
	var slen, start, end, space int
	var lastLine bool
	var result []string
	for i := 0; i < len(input); i++ {

		slen += len(input[i])
	
		if slen+i-start > long || i == len(input)-1 {

			if slen+i-start > long {

				slen -= len(input[i])
				end = i - 1
				lastLine = false
			} else {
				end = i
				lastLine = true
			}

			space = long - slen

			var mspace int
			var extra int
			if lastLine {
				mspace = 1
				extra = 0
			} else {
				mspace = space / (end - start)
				extra = space - mspace*(end-start)
			}

			var line string = input[start]

			for j := start + 1; j <= i-1; j++ {
				for k := 0; k < mspace && space > 0; k++ {
					line += " "
					space--
				}
				if j-start-1 < extra {
					line += " "
					space--
				}
				line += input[j]
			}

			if space > 0 {
				for ; space > 0; space-- {
					line += " "
				}
			}
			result = append(result, line)
			start = i
			i = end
			slen = 0
		}

	}

	return result

}

func main() {
	test := []string{"This", "is", "an", "example", "of", "text", "justification."}
	res := fullJustify(test, 16)
	for _, v := range res {
		fmt.Println(v)
	}
}
