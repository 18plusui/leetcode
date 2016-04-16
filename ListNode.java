package leetcode;

import java.util.Random;

//  Definition for singly-linked list.
public class ListNode {
	int val;
	ListNode next;

	ListNode(int x) {
		val = x;
		next = null;
	}

	public ListNode createListNode(int i, ListNode ln) {

		Random r = new Random();
		int val = r.nextInt(10);

		ln.next = new ListNode(val);

		if (i == 0) {
			return ln;
		}
		i--;
		createListNode(i, ln.next);
		return ln.next;
	}

	public static void main(String[] args) {

		ListNode listnode = new ListNode(0);

		ListNode result = listnode.createListNode(10, listnode);

		System.out.println(result.next.val + " : " + result.val);

	}
}
