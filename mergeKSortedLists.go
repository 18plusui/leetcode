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

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = nil
	curr := dummy

	length := len(lists)
	for {
		minPos := -1
		for i := 0; i < length; i++ {
			if lists[i] != nil && (minPos == -1 || lists[minPos].Val > lists[i].Val) {
				minPos = i
			}
		}
		if minPos == -1 {
			break
		}
		curr.Next = lists[minPos]
		lists[minPos] = lists[minPos].Next
		curr = curr.Next
	}

	return dummy.Next
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

	var lns []*ListNode
	lns = append(lns, la)
	lns = append(lns, lb)

	ln := mergeKLists(lns)
	printList(ln)
}
