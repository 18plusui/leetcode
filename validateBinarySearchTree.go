// Source : https://oj.leetcode.com/problems/validate-binary-search-tree/
// Author : 18plusui
// Date   : 2016-03-07

/********************************************************************************** 
* 
* Given a binary tree, determine if it is a valid binary search tree (BST).
* 
* Assume a BST is defined as follows:
* 
* The left subtree of a node contains only nodes with keys less than the node's key.
* The right subtree of a node contains only nodes with keys greater than the node's key.
* Both the left and right subtrees must also be binary search trees.
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

func CreateTree(inputArr []int) *Tree {
	var t *Tree
	for _, v := range inputArr {

		t = insert(t, v)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v < t.Value {

		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)
	return t
}

func isValidBST(t *Tree) bool {

	fmt.Println(t.Value)

	if t.Left != nil && t.Right != nil {
		if t.Value < t.Left.Value || t.Value > t.Right.Value {

			return false
		}
		isValidBST(t.Left)
		isValidBST(t.Right)
	} else if t.Right != nil {
		if t.Value > t.Right.Value {

			return false
		}
		isValidBST(t.Right)
	} else if t.Left != nil {
		if t.Value < t.Left.Value {

			return false
		}
		isValidBST(t.Left)
	}

	return true
}

func main() {
	inputArr := []int{6, 4, 5, 1, 7, 8}
	t1 := CreateTree(inputArr)
	fmt.Println(isValidBST(t1))
	fmt.Println(t1.Left.Right.Value)

}
