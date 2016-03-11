// Source : https://oj.leetcode.com/problems/triangle/
// Author : 18plusui
// Date   : 2016-03-11

/**********************************************************************************
*
* Given a triangle, find the minimum path sum from top to bottom.
* Each step you may move to adjacent numbers on the row below.
*
* For example, given the following triangle
*
* [
*      [2],
*     [3,4],
*    [6,5,7],
*   [4,1,8,3]
* ]
*
* The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
*
* Note:
* Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
*
**********************************************************************************/

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func minTotal(inputTri [][]int) int {
	var minT int
	for i := 0; i < len(inputTri); i++ {
		if i == 0 {
			minT += inputTri[0][0]
			continue
		}

		intSort := inputTri[i]
		sort.Ints(intSort)
		minT += intSort[0]

	}
	return minT
}

func createTriangle(size int) [][]int {
	var row int

	for i := 1; i <= size; i++ {

		if (i*(i+1))/2 == size {
			row = i
			break
		}

	}

	res := make([][]int, row)
	for i := 0; i < row; i++ {
		innerLen := i + 1
		res[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			res[i][j] = r1.Intn(30)
			time.Sleep(1)
		}
	}

	fmt.Println(res)
	return res
}

func main() {
	inputTri := createTriangle(10)
	fmt.Println(minTotal(inputTri))
}
