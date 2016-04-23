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

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {

	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	list1 := []int{1, 2, 3, 4, 6}
	list2 := []int{1, 2, 3, 7, 6}
	pq := make(PriorityQueue, len(list1))
	i := 0
	for priority, value := range list1 {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	for priority, value := range list2 {

		item := &Item{
			value:    value,
			priority: priority,
		}
		heap.Push(&pq, item)
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Print(item.value)
	}
}

