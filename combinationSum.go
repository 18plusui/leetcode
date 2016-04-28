// Source : https://oj.leetcode.com/problems/combination-sum/
// Author : 18plusui
// Date   : 2016-04-28

/**********************************************************************************
*
* Given a set of candidate numbers (C) and a target number (T), find all unique combinations
* in C where the candidate numbers sums to T.
*
* The same repeated number may be chosen from C unlimited number of times.
*
* Note:
*
* All numbers (including target) will be positive integers.
* Elements in a combination (a1, a2, … , ak) must be in non-descending order. (ie, a1 ≤ a2 ≤ … ≤ ak).
* The solution set must not contain duplicate combinations.
*
* For example, given candidate set 2,3,6,7 and target 7,
* A solution set is:
* [7]
* [2, 2, 3]
*
**********************************************************************************/
package main

import (
	"fmt"
	"sort"
)

func combinationSumHelper(candidates []int, start, target int, solution []int, result [][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		result = append(result, solution)
		fmt.Println(solution)
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		solution = append(solution, candidates[i])
		combinationSumHelper(candidates, i, target-candidates[i], solution, result)

		solution = solution[:len(solution)-1]

	}

}

func combinationSum(candidates []int, target int) {
	var result [][]int
	if len(candidates) <= 0 {
		return
	}

	sort.IntsAreSorted(candidates)
	var solution []int
	combinationSumHelper(candidates, 0, target, solution, result)

}

func main() {
	combinationSum([]int{1, 3, 2, 4, 5, 6}, 7)
}
