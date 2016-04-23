// Source : https://oj.leetcode.com/problems/merge-k-sorted-lists/
// Author : 18plusui
// Date   : 2016-04-23

/**********************************************************************************
*
* Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*
***********************************************************************************/

func mergeKLists(lists []*ListNode) *ListNode {
    dummy := new(ListNode)
    dummy.Next = nil
    curr := dummy
    
    length := len(lists)
    for ; ; {
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
