// Source : https://oj.leetcode.com/problems/swap-nodes-in-pairs/
// Author : 18plusui
// Date   : 2016-04-24

/**********************************************************************************
*
* Given a linked list, swap every two adjacent nodes and return its head.
*
* For example,
* Given 1->2->3->4, you should return the list as 2->1->4->3.
*
* Your algorithm should use only constant space. You may not modify the values in the list,
* only nodes itself can be changed.
*
**********************************************************************************/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) setVal(val int) *ListNode {
	l.Val = val
	l.Next = nil
	return l
}

func swapPairs(head *ListNode) *ListNode {
	for p := head; p != nil && p.Next != nil; p = p.Next.Next {
		p.Val, p.Next.Val = p.Next.Val, p.Val
	}
	return head
}

func createList(intSlice []int, ln *ListNode) *ListNode {
	n := len(intSlice)
	head := new(ListNode)

	if n != 0 {
		val := intSlice[n-1]
		ln.Next = head.setVal(val)
	} else {
		return ln
	}

	n--

	createList(intSlice[:n], ln.Next)
	return ln.Next
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Println("===", list.Val)
		list = list.Next

	}
}

func main() {
	a := []int{4, 4, 3, 2, 1}

	la := new(ListNode)
	la = createList(a, la)
	ln := swapPairs(la)
	printList(ln)
}
