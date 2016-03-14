// Source : https://leetcode.com/problems/the-skyline-problem/
// Author : 18plusui
// Date   : 2016-03-13

/**********************************************************************************
 *
 * A city's skyline is the outer contour of the silhouette formed by all the buildings
 * in that city when viewed from a distance. Now suppose you are given the locations and
 * height of all the buildings as shown on a cityscape photo (Figure A), write a program
 * to output the skyline formed by these buildings collectively (Figure B).
 *
 *  ^                                        ^
 *  |                                        |
 *  |                                        |
 *  |    +-----+                             |    O-----+
 *  |    |     |                             |    |     |
 *  |    |     |                             |    |     |
 *  |    |  +--+------+                      |    |     O------+
 *  |    |  |         |                      |    |            |
 *  |  +-+--+----+    |   +------+           |  O-+            |   O------+
 *  |  |         |    |   |      |           |  |              |   |      |
 *  |  |         |    |   |    +-+--+        |  |              |   |      O--+
 *  |  |         |    |   |    |    |        |  |              |   |         |
 *  |  |         |    |   |    |    |        |  |              |   |         |
 *  |  |         |    |   |    |    |        |  |              |   |         |
 *  |  |         |    |   |    |    |        |  |              |   |         |
 *  +--+---------+----+---+----+----+--->    +--+--------------O---+---------O--->
 *
 *   https://leetcode.com/static/images/problemset/skyline1.jpg
 *   https://leetcode.com/static/images/problemset/skyline2.jpg
 *
 * The geometric information of each building is represented by a triplet of integers [Li, Ri, Hi],
 * where Li and Ri are the x coordinates of the left and right edge of the ith building, respectively,
 * and Hi is its height. It is guaranteed that 0 ≤ Li, Ri ≤ INT_MAX, 0 , and Ri - Li > 0.
 * You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.
 *
 * For instance, the dimensions of all buildings in Figure A are recorded as:
 *  [ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] .
 *
 * The output is a list of "key points" (red dots in Figure B) in the format of
 * [ [x1,y1], [x2, y2], [x3, y3], ... ] that uniquely defines a skyline.
 * A key point is the left endpoint of a horizontal line segment.
 *
 * Note that the last key point, where the rightmost building ends, is merely used to mark
 * the termination of the skyline, and always has zero height. Also, the ground in between
 * any two adjacent buildings should be considered part of the skyline contour.
 *
 * For instance, the skyline in Figure B should be represented as:
 *  [ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ].
 *
 * Notes:
 *
 *  - The number of buildings in any input list is guaranteed to be in the range [0, 10000].
 *  - The input list is already sorted in ascending order by the left x position Li.
 *  - The output list must be sorted by the x position.
 *  - There must be no consecutive horizontal lines of equal height in the output skyline.
 *    For instance, [...[2 3], [4 5], [7 5], [11 5], [12 7]...] is not acceptable;
 *    the three lines of height 5 should be merged into one in the final output as such:
 *    [...[2 3], [4 5], [12 7], ...]
 *
 * Credits: Special thanks to @stellari for adding this problem,
 *          creating these two awesome images and all test cases.
 *
 **********************************************************************************/
package main

import (
	"fmt"
	"sort"
)

type point struct {
	x int
	y int
}

type arrPoint []point

func (s arrPoint) Len() int      { return len(s) }
func (s arrPoint) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByX struct{ arrPoint }

func (s ByX) Less(i, j int) bool {

	if s.arrPoint[i].x == s.arrPoint[j].x && s.arrPoint[i].y <= s.arrPoint[j].y {
		return true
	}
	return s.arrPoint[i].x < s.arrPoint[j].x
}

func getSkyline(input [][]int) []point {
	var edges []point

	var left, right, height int
	for i := 0; i < len(input); i++ {
		left = input[i][0]
		right = input[i][1]
		height = input[i][2]

		edges = append(edges, point{left, -height})
		edges = append(edges, point{right, height})

	}

	sort.Sort(ByX{edges})

	fmt.Println(edges)
	var res []point

	var pre, curr int
	var m []int
	var mLen int
	m = append(m, 0)

	for i := 0; i < len(edges); i++ {
		e := edges[i]

		if e.y < 0 {
			if pre > -e.y {
				m = append(m, -e.y)
				continue
			}
			m = append(m, -e.y)

		} else {
			for j, v := range m {
				if v == e.y {

					m = append(m[:j], m[j+1:]...)

				}

			}

		}
		mLen = len(m)

		curr = m[mLen-1]

		if curr != pre {
			res = append(res, point{e.x, curr})
			pre = curr
		}

	}

	return res
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {

	res := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}

	fmt.Println(res)
	// getSkyline(res)
	fmt.Println(getSkyline(res))
}
