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
