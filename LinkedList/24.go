package LinkedList

/*
   题目描述：
	Middle 两两交换链表中的结点
	给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/

/*
   解题思路：
	递归
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	q := head.Next.Next
	head.Next.Next = head
	head.Next = swapPairs(q)
	return p
}
