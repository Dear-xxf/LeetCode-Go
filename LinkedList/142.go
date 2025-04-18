package LinkedList

/*
	题目描述：
	Middle 环形链表2
	给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
	为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
	如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
	不允许修改 链表。
*/

/*
	解题思路：
	与环形链表1一样，哈希表的做法很好想，但是双指针的进阶版就不是很容易了
	记录力扣的官方题解
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
