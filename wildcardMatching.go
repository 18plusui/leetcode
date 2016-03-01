package main

import (
	"fmt"
)

func isMatch(inputStr string, matchStr string) bool {
	inputStart := 0
	matchStart := 0

	for inputStart != len(inputStr) {
		if matchStr[matchStart:matchStart+1] == "*" {
			matchStart++
			if matchStart == len(matchStr) {
				return true
			}

		} else if matchStr[matchStart:matchStart+1] == "?" ||
			inputStr[inputStart:inputStart+1] == matchStr[matchStart:matchStart+1] {
			matchStart++
			inputStart++
			if matchStart == len(matchStr) && matchStart == len(inputStr) {
				return true

			}
		} else if inputStart != 0 {
			inputStart++

		}

	}

	return false

}

func main() {

	fmt.Println(isMatch("hello", "h??o"))
}
