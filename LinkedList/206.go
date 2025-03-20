package LinkedList

/*
	题目描述：
	Easy 反转链表
	给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

/*
	解题思路：
	递归
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
