package LinkedList

/*
	题目描述：
	Easy 相交链表
	给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

/*
	解题思路：
	1.使用哈希表，遍历一个链表之后，记录所有的val在哈希表中。遍历第二个链表时，如果存在重复，即返回重复的值
	2.双指针法，两个指针同时向后遍历，当两个指针相遇时即为要找的结点。当一个链表上的指针走到最后时，如果不是两个指针同时为空
	  则将这个指针赋到另外一个链表的头指针上。本质上是令两个链表的长度一致
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1.哈希表
	//nodeHash := map[*ListNode]int{}
	//p, q := headA, headB
	//for p != nil {
	//	nodeHash[p] = p.Val
	//	p = p.Next
	//}
	//for q != nil {
	//	if nodeHash[q] == q.Val {
	//		return q
	//	}
	//	q = q.Next
	//}
	//return nil

	// 2. 双指针
	p, q := headA, headB
	for p != q {
		p = p.Next
		q = q.Next
		if p == nil && q == nil { // 如果同时为空
			return nil
		} else if p == nil {
			p = headB
		} else if q == nil {
			q = headA
		}
	}
	return p
}
