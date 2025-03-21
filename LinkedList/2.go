package LinkedList

/*
   题目描述：
	Middle 两数相加
	给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
	请你将两个数相加，并以相同形式返回一个表示和的链表。
	你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

/*
   解题思路：
	建立一个临时变量，记录每一位相加后溢出10的数字，计入后面的位数计算
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := int((l1.Val + l2.Val) / 10)
	res := &ListNode{Val: (l1.Val + l2.Val) % 10, Next: nil}
	head := res
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil || l2 != nil {
		var cur int
		if l1 == nil {
			cur = l2.Val + temp
			l2 = l2.Next
		} else if l2 == nil {
			cur = l1.Val + temp
			l1 = l1.Next
		} else {
			cur = l1.Val + l2.Val + temp
			l1 = l1.Next
			l2 = l2.Next
		}
		temp = int(cur / 10)
		res.Next = &ListNode{Val: cur % 10, Next: nil}
		res = res.Next
	}
	if temp != 0 {
		res.Next = &ListNode{Val: temp, Next: nil}
	}
	return head
}
