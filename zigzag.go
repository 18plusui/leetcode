// Source : https://oj.leetcode.com/problems/zigzag-conversion/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: 
* (you may want to display this pattern in a fixed font for better legibility)
* 
* P   A   H   N
* A P L S I I G
* Y   I   R
* 
* And then read line by line: "PAHNAPLSIIGYIR"
* 
* Write the code that will take a string and make this conversion given a number of rows:
* 
* string convert(string text, int nRows);
* 
* convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
* 
*               
**********************************************************************************/
package main

import (
	"fmt"
	"strings"
)

var x int = 0
var step int = 0
var inputArr []string

func zigzag(inputStr string, nRows int) string {
	if len(inputStr) < nRows || nRows <= 0 {
		return "you maybe not understand rule !!"
	}

	inputArr = strings.Split(inputStr, "")
	tmpArr := make([]string, nRows)

	for i := 0; i < len(inputArr); i++ {

		if x == nRows-1 {
			step = -1

		}

		if x == 0 {
			step = 1
		}

		fmt.Println("row number is : ", x)

		tmpArr[x] = tmpArr[x] + inputArr[i]
		x += step

		fmt.Println(tmpArr)

	}

	return ""
}

func main() {
	zigzag("helloworld", 4)
}
