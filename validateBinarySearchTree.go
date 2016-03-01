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
