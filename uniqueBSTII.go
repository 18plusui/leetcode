// Source : https://oj.leetcode.com/problems/unique-binary-search-trees-ii/
// Author : 18plusui
// Date   : 2016-03-08

/**********************************************************************************
*
* Given n, generate all structurally unique BST's (binary search trees) that store values 1...n.
*
* For example,
* Given n = 3, your program should return all 5 unique BST's shown below.
*
*    1         3     3      2      1
*     \       /     /      / \      \
*      3     2     1      1   3      2
*
* confused what "{1,#,2,3}" means? > read more on how binary tree is serialized on OJ.
**********************************************************************************/
package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) newTree(n int) *Tree {
	t.Value = n
	t.Left = nil
	t.Right = nil
	return t
}

func generateTreeI(n int) []*Tree {
	var resArr []*Tree
	resArr = generateTree(1, n)
	return resArr
}

func generateTree(low, high int) []*Tree {
	var res []*Tree

	if low > high || low <= 0 || high <= 0 {
		res = append(res, nil)
		return res
	}

	if low == high {
		tree := new(Tree)
		res = append(res, tree.newTree(low))
		return res
	}

	for i := low; i <= high; i++ {
		var arrL []*Tree = generateTree(low, i-1)
		var arrR []*Tree = generateTree(i+1, high)
		for l := 0; l < len(arrL); l++ {
			for r := 0; r < len(arrR); r++ {
				root := new(Tree)
				root.newTree(i)
				root.Left = arrL[l]
				root.Right = arrR[r]
				res = append(res, root)
			}
		}

	}
	return res
}

func printer(t *Tree) {
	if t == nil {
		fmt.Println("#")
		return

	}
	fmt.Println(t.Value)
	printer(t.Left)
	printer(t.Right)

}

func main() {
	t := new(Tree)
	fmt.Println(t.newTree(0))
	var res []*Tree = generateTreeI(3)
	for i, v := range res {
		fmt.Println("the loop number : ", i)
		printer(v)
	}

}
