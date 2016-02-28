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
