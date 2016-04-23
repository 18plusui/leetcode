// Source : https://oj.leetcode.com/problems/merge-k-sorted-lists/
// Author : 18plusui
// Date   : 2016-04-23

/**********************************************************************************
*
* Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*
***********************************************************************************/

package main

import (
	"fmt"
)

func mergeTwoList(l1, l2, result *ListNode) *ListNode {

	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			result.Next = l1
			l1 = l1.Next
			if l1 == nil {
				result.Next.Next = l2
			}

		} else {
			result.Next = l2
			l2 = l2.Next
			if l2 == nil {
				result.Next.Next = l1
			}
		}

		mergeTwoList(l1, l2, result.Next)
	}

	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) setVal(val int) *ListNode {
	l.Val = val
	l.Next = nil
	return l
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
	a := []int{3, 3, 3, 2, 1}

	la := new(ListNode)
	la = createList(a, la)

	b := []int{5, 4, 3, 2, 1}

	lb := new(ListNode)
	lb = createList(b, lb)

	result := new(ListNode)
	ln := mergeTwoList(la, lb, result)

	printList(ln)
}
