package LinkedList

/*
   题目描述：
	Easy 环形链表
	给你一个链表的头节点 head ，判断链表中是否有环。
	如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
	为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
	注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
	如果链表中存在环 ，则返回 true 。 否则，返回 false 。
*/

/*
	解题思路：
	哈希表很容易想到。这里实现快慢指针
*/

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		// 快指针每次前进两步 慢指针每次前进一步
		// 如果存在环，两个指针一定会相遇
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
