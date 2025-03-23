package LinkedList

/*
   题目描述：
	Hard 合并K个升序链表
	给你一个链表数组，每个链表都已经按升序排列。
	请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

/*
   解题思路：
	分治法。递归的思路，每一次将数组从中间分开；
	最后只剩两个数组的时候，使用Q21的合并两个升序链表的函数
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	pivot := len(lists) / 2
	left := mergeKLists(lists[:pivot])
	right := mergeKLists(lists[pivot:])
	return mergeTwoLists(left, right)
}
