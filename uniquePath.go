// Source : https://oj.leetcode.com/problems/unique-paths/
// Author : 18plusui
// Date   : 2016-03-07

/**********************************************************************************
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot is trying to reach
 * the bottom-right corner of the grid (marked 'Finish' in the diagram below).
 *
 *
 *    start
 *    +---------+----+----+----+----+----+
 *    |----|    |    |    |    |    |    |
 *    |----|    |    |    |    |    |    |
 *    +----------------------------------+
 *    |    |    |    |    |    |    |    |
 *    |    |    |    |    |    |    |    |
 *    +----------------------------------+
 *    |    |    |    |    |    |    |----|
 *    |    |    |    |    |    |    |----|
 *    +----+----+----+----+----+---------+
 *                                   finish
 *
 *
 * How many possible unique paths are there?
 *
 * Above is a 3 x 7 grid. How many possible unique paths are there?
 *
 * Note: m and n will be at most 100.
 *
 **********************************************************************************/

package main

import (
	"fmt"
)

func uniquePath(m, n int) int {
	//create 2d arrays
	arr2d := make([][]int, m)

	for i := 0; i < len(arr2d); i++ {
		arr2d[i] = make([]int, n)
	}

	var MaxNum int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				arr2d[i][j] = 0
				MaxNum = arr2d[i][j]
			} else {

				if i > 0 && j > 0 {
					arr2d[i][j] = arr2d[i-1][j] + arr2d[i][j-1]
					MaxNum = arr2d[i][j]
				} else {

					arr2d[i][j] = 1
					MaxNum = arr2d[i][j]

				}
			}

		}
	}
	fmt.Println(arr2d[0])
	fmt.Println(arr2d[1])
	fmt.Println(arr2d[2])

	return MaxNum
}

func main() {
	fmt.Print(uniquePath(3, 7))

}
