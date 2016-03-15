// Source : https://oj.leetcode.com/problems/symmetric-tree/
// Author : 18plusui
// Date   : 2016-03-15

/**********************************************************************************
*
* Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
*
* For example, this binary tree is symmetric:
*
*     1
*    / \
*   2   2
*  / \ / \
* 3  4 4  3
*
* But the following is not:
*
*     1
*    / \
*   2   2
*    \   \
*    3    3
*
* Note:
* Bonus points if you could solve it both recursively and iteratively.
*
* confused what "{1,#,2,3}" means? > read more on how binary tree is serialized on OJ.
*
* OJ's Binary Tree Serialization:
*
* The serialization of a binary tree follows a level order traversal, where '#' signifies
* a path terminator where no node exists below.
*
* Here's an example:
*
*    1
*   / \
*  2   3
*     /
*    4
*     \
*      5
*
* The above binary tree is serialized as "{1,2,3,#,#,4,#,#,5}".
*
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

func CreateSymmetricTree(inputArr []int) *Tree {
	var t *Tree
	var count int
	for _, v := range inputArr {

		t = insertLeft(t, v, count)
	}

	count = 0
	for _, v := range inputArr {

		if v == t.Value {
			continue
		}

		t = insertRight(t, v, count)
	}
	return t
}

func insertLeft(t *Tree, v, count int) *Tree {
	count++
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if t.Left == nil {

		t.Left = insertLeft(t.Left, v, count)
		return t
	}

	if t.Left == nil || count == 1 {

		t.Left = insertLeft(t.Left, v, count)
		return t
	}

	t.Right = insertLeft(t.Right, v, count)

	return t
}

func insertRight(t *Tree, v, count int) *Tree {
	count++
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if t.Right == nil {

		t.Right = insertRight(t.Right, v, count)
		return t
	}

	if t.Right == nil || count == 1 {

		t.Right = insertRight(t.Right, v, count)
		return t
	}

	t.Left = insertRight(t.Left, v, count)

	return t
}

// above code is generate Symmetric Tree

func isSymmetric(t *Tree) bool {
	if t == nil {
		return true
	}

	return isSymmetricNode(t.Left, t.Right)
}

func isSymmetricNode(left, right *Tree) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return (left.Value == right.Value) && isSymmetricNode(left.Left, right.Right) && isSymmetricNode(left.Right, right.Left)
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

	inputArr := []int{6, 5, 4, 3, 2, 1}
	t := CreateSymmetricTree(inputArr)

	printer(t)
	fmt.Println(isSymmetric(t))

}
