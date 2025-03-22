package LinkedList

/*
   题目描述：
	Hard K个一组翻转链表
	给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

/*
   解题思路：
	不算困难的Hard题。用递归可以解决，reverseList使用了206中的函数
	优化点在于两个函数都是用了递归，空间开销较大
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	fakeHead := &ListNode{Next: head}
	tail := fakeHead
	for i := 0; i < k; i++ {
		tail = tail.Next
		if tail == nil {
			return fakeHead.Next
		}
	}
	nex := tail.Next
	tail.Next = nil
	previousHead := head
	head = reverseList(head)
	nex = reverseKGroup(nex, k)
	previousHead.Next = nex
	return tail
}
