package main

import (
	"fmt"
	"strings"
)

var boardStr []string
var arr1dlen int
var resVal bool

func exist(board [][]string, word string) bool {
	if len(board) <= 0 || len(word) <= 0 {
		return false
	}

	mask := make([][]int, len(board))
	for i := 0; i < 3; i++ {
		mask[i] = make([]int, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			mask[i][j] = 0
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			wordArr := strings.Split(word, "")

			if board[i][j] == wordArr[0] {
				if getExist(board, wordArr, 0, i, j, mask) {
					fmt.Println(mask)
					return true
				}
			}

		}
	}

	return false
}

func getExist(board [][]string, word []string, index int, row int, col int, mask [][]int) bool {

	i := row
	j := col

	if board[i][j] == word[index] && mask[i][j] == 0 {

		mask[i][j] = 1
		// fmt.Println("this is index length :", index, i, j)

		if index+1 == len(word) {
			return true
		}
		index++

		if i+1 < len(board) && getExist(board, word, index, i+1, j, mask) ||
			(i > 0 && getExist(board, word, index, i-1, j, mask)) ||
			(j+1 < len(board[i]) && getExist(board, word, index, i, j+1, mask)) ||
			(j > 0 && getExist(board, word, index, i, j-1, mask)) {
			return true
		}
		mask[i][j] = 0
	}

	return false
}

func main() {

	boardStr = []string{"ABCE", "SFCS", "ADEE"}
	arr2dLen := len(boardStr)

	board := make([][]string, arr2dLen)
	for i := 0; i < arr2dLen; i++ {
		arr1dlen = len(boardStr[i])
		board[i] = make([]string, arr1dlen)
		board[i] = strings.Split(boardStr[i], "")
	}

	word := "ABCCED"
	fmt.Println(exist(board, word))

}
