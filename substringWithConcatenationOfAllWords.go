// Source : https://oj.leetcode.com/problems/substring-with-concatenation-of-all-words/
// Author : 18plusui
// Date   : 2016-04-27

/**********************************************************************************
*
* You are given a string, S, and a list of words, L, that are all of the same length.
* Find all starting indices of substring(s) in S that is a concatenation of each word
* in L exactly once and without any intervening characters.
*
* For example, given:
* S: "barfoothefoobarman"
* L: ["foo", "bar"]
*
* You should return the indices: [0,9].
* (order does not matter).
*
**********************************************************************************/

package main

import (
	"fmt"
)

func findSubstr(S string, L []string) []int {
	var result []int

	if len(S) <= 0 || len(L) <= 0 {
		return result
	}

	n, m, l := len(S), len(L), len(L[0])

	expected := make(map[string]int)
	for i := 0; i < m; i++ {
		_, ok := expected[L[i]]
		if ok {
			expected[L[i]]++
		} else {
			expected[L[i]] = 1
		}
	}

	for i := 0; i < l; i++ {
		actual := make(map[string]int)
		count, winleft := 0, i

		for j := i; j <= n-l; j += l {

			word := S[j : j+l]
			_, ok := expected[word]
			if !ok {
				for k := range actual {
					delete(actual, k)
				}
				count = 0
				winleft = j + l
				continue
			}
			count++

			_, ok = actual[word]
			if !ok {
				actual[word] = 1
			} else {
				actual[word]++
			}

			if actual[word] > expected[word] {

				tmp := S[winleft : winleft+l]
				count--
				actual[tmp]--
				winleft += l
				for tmp != word {
					tmp = S[winleft : winleft+l]
					count--
					actual[tmp]--
					winleft += l
				}

			}

			if count == m {
				result = append(result, winleft)
				tmp := S[winleft : winleft+l]
				actual[tmp]--
				winleft += l
				count--
			}

		}
	}

	return result

}

func main() {

	fmt.Println(findSubstr("barfoothefoobarman", []string{"foo", "bar"}))

}
