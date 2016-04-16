package leetcode;

public class MergeTowSortList {

	public ListNode merge2List(ListNode a, ListNode b) {

		ListNode result = new ListNode(0);
		ListNode h = result;
		while (a != null && b != null) {
			if (a.val < b.val) {
				result.next = a;
				a = a.next;
			} else {
				result.next = b;
				b = b.next;
			}
			result = result.next;

		}

		if (a != null)
			result.next = a;
		else
			result.next = b;

		return h;

	}

	public static void main(String[] args) {
		ListNode ln = new ListNode(0);

		ListNode l1 = ln.createListNode(10, ln);
		ListNode l2 = ln.createListNode(10, ln);

		MergeTowSortList mtsl = new MergeTowSortList();

		ListNode rst = mtsl.merge2List(l1.next, l2.next);

		while (rst != null) {
			System.out.println("rst : " + rst.val);
			rst = rst.next;

		}
	}

}
