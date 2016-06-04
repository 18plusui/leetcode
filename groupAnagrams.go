// Source : https://oj.leetcode.com/problems/anagrams/
// Author : 18plusui
// Date   : 2016-06-04

/**********************************************************************************
 *
 * Given an array of strings, group anagrams together.
 *
 * For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Return:
 *
 * [
 *   ["ate", "eat","tea"],
 *   ["nat","tan"],
 *   ["bat"]
 * ]
 *
 * Note:
 *
 * For the return value, each inner list's elements must follow the lexicographic order.
 * All inputs will be in lower-case.
 *
 **********************************************************************************/

package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {

	strlen := len(strs)
	result := make([][]string, strlen)
	m := make(map[string]int)
	var i, j int

	for i, j = 0, 0; i < strlen; i++ {
		word := strs[i]
		sortword := SortString(word)

		if val, ok := m[sortword]; ok {
			result[val] = append(result[val], word)
		} else {
			m[sortword] = j
			result[j] = append(result[j], word)
			j++
		}

	}

	return result[0:j]
}

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}
