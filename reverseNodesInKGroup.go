// Source : https://oj.leetcode.com/problems/reverse-nodes-in-k-group/
// Author : 18plusui
// Date   : 2016-04-24

/**********************************************************************************
*
* Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
*
* If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
*
* You may not alter the values in the nodes, only nodes itself may be changed.
*
* Only constant memory is allowed.
*
* For example,
* Given this linked list: 1->2->3->4->5
*
* For k = 2, you should return: 2->1->4->3->5
*
* For k = 3, you should return: 3->2->1->4->5
*
**********************************************************************************/
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	lists := getListSlice(head)

	if len(lists) < k {
		return head
	}

	for n := 1; n <= len(lists)/k; n++ {

		for i := 0; i+(n-1)*k != n*k-i-1; i++ {

			lists[i+(n-1)*k].Val, lists[n*k-i-1].Val = lists[n*k-i-1].Val, lists[i+(n-1)*k].Val

		}
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

func getListSlice(list *ListNode) []*ListNode {

	var ListSlice []*ListNode

	for list != nil {
		ListSlice = append(ListSlice, list)
		list = list.Next

	}

	return ListSlice
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next

	}
}

func main() {
	a := []int{4, 4, 3, 2, 1, 7, 8, 5, 6, 1, 2}

	la := new(ListNode)
	la = createList(a, la)
	// printList(la)

	ln := reverseKGroup(la, 5)
	printList(ln)
}
