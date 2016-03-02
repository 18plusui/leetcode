// Source : https://oj.leetcode.com/problems/valid-sudoku/
// Author : 18plusui
// Date   : 2016-03-02

/**********************************************************************************
*
* Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.
*
* The Sudoku board could be partially filled, where empty cells are filled with
* the character '.'.
* A partially filled sudoku which is valid.
*
* Note:
* > A valid Sudoku board (partially filled) is not necessarily solvable.
*   Only the filled cells need to be validated.
*
**********************************************************************************/

package main

import (
	"fmt"
)

const (
	cnt int = 9
)

func isValidSudoku(inputArr [9][9]int) bool {

	var row_mask [cnt][cnt]bool
	var col_mask [cnt][cnt]bool
	var partArea_mask [cnt][cnt]bool

	for r := 0; r < len(inputArr); r++ {
		for c := 0; c < len(inputArr[r]); c++ {

			if inputArr[r][c] == 0 {
				continue
			}
			idx := inputArr[r][c] - 1

			if row_mask[r][idx] == true {
				return true
			}
			row_mask[r][idx] = true

			if col_mask[c][idx] == true {
				return true
			}
			row_mask[c][idx] = true

			area := (r/3)*3 + (c / 3)
			if partArea_mask[area][idx] == true {
				return false
			}
			partArea_mask[area][idx] = true
		}
	}
	return true
}

func main() {
	inputArr := [9][9]int{
		{0, 0, 6, 4, 8, 1, 3, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 4, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 9},
		{8, 0, 0, 0, 9, 0, 0, 0, 4},
		{6, 0, 0, 3, 4, 2, 0, 0, 1},
		{5, 0, 0, 0, 6, 0, 0, 0, 2},
		{3, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 9, 0, 0, 0, 0, 0, 7, 0},
		{0, 0, 5, 7, 1, 6, 2, 0, 0},
	}

	fmt.Println(isValidSudoku(inputArr))
}
