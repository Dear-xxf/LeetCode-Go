package LinkedList

/*
   题目描述：
	Middle 删除链表的倒数第N个结点
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

/*
   解题思路：
	双指针。一个指针先前进n步，之后第二个指针再前进
	可以找到倒数第n个结点
	需要注意的特殊情况是删除的正好是头结点这个case
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head} // 使用虚拟头节点，防止删除头节点时出错
	p1, p2 := dummyHead, dummyHead
	count := 0
	for p1 != nil {
		p1 = p1.Next
		if count >= n+1 {
			p2 = p2.Next
		}
		count++
	}
	p2.Next = p2.Next.Next
	return dummyHead.Next
}
