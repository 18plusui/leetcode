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
