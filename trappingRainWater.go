// Source : https://oj.leetcode.com/problems/trapping-rain-water/
// Author : 18plusui
// Date   : 2016-03-12

/**********************************************************************************
*
* Given n non-negative integers representing an elevation map where the width of each bar is 1,
* compute how much water it is able to trap after raining.
*
* For example,
* Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
*
*     ^
*     |
*   3 |                       +--+
*     |                       |  |
*   2 |          +--+xxxxxxxxx|  +--+xx+--+
*     |          |  |xxxxxxxxx|  |  |xx|  |
*   1 |   +--+xxx|  +--+xxx+--+  |  +--+  +--+
*     |   |  |xxx|  |  |xxx|  |  |  |  |  |  |
*   0 +---+--+---+--+--+---+--+--+--+--+--+--+----->
*       0  1   0  2  1   0  1  3  2  1  2  1
*
* The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case,
* 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
*
**********************************************************************************/

package main

import (
	"fmt"
)

func trappingRain(input []int) int {
	var res, maxHigh, maxIdx, preHigh int

	for i, v := range input {
		if v > maxHigh {
			maxHigh = v
			maxIdx = i
		}
	}

	for i := 0; i < maxIdx; i++ {
		if input[i] > preHigh {
			preHigh = input[i]
		}
		res += preHigh - input[i]
	}

	preHigh = 0
	for i := len(input) - 1; i > maxIdx; i-- {
		if input[i] > preHigh {
			preHigh = input[i]
		}
		res += preHigh - input[i]
	}
	return res
}

func main() {

	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trappingRain(input))
}
