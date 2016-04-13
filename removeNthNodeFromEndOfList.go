// Source : https://oj.leetcode.com/problems/remove-nth-node-from-end-of-list/
// Author : 18plusui
// Date   : 2016-04-14

/********************************************************************************** 
* 
* Given a linked list, remove the nth node from the end of list and return its head.
* 
* For example,
* 
*    Given linked list: 1->2->3->4->5, and n = 2.
* 
*    After removing the second node from the end, the linked list becomes 1->2->3->5.
* 
* Note:
* Given n will always be valid.
* Try to do this in one pass.
*               
**********************************************************************************/

package main

import (
	"fmt"
)

type ListNode struct {
     Val int
     Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
   
   if head == nil || n <= 0 {
   	   return nil
   }
   
   var fakeHead ListNode = ListNode{0,head}

   head = &fakeHead

   p1 := new(ListNode)
   p2 := new(ListNode)

   p1=head
   p2=head

   for i := 0; i < n; i++ {
   	   if p2 == nil {
   	   	  return nil
   	   }
   	   p2 = p2.Next
   }

   for p2.Next != nil{
   	   p2 = p2.Next
   	   p1 = p1.Next
   }

   p1.Next = p1.Next.Next

   return head.Next
}

// func makeListNode(head *ListNode,num []int) *ListNode{
    
//     numlen := len(num)
    
//     if numlen != 0 {
//     	head.Val = num[0]
//     	makeListNode(head.Next, num[1:])	
//     }
    
//     return head

// }

func main() {
	
}
