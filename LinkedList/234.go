package LinkedList

/*
	题目描述：
	Easy 回文链表
	给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
*/

/*
	解题思路：
	初步想法是将链表复制，复制后的链表整个反转，再逐一对比；可以用快慢指针来优化，只反转后半部分
	代码记录了官方递归题解，比较巧妙，再写了一个递归函数，利用递归特性模拟双指针
	每一次递归相当于后面指针向前移动
	frontPointer = frontPointer.Next 则控制每次判断之后前面的指针向后移动
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}
