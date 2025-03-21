package LinkedList

/*
   题目描述：
	Easy 合并两个有序链表
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

/*
   解题思路：
	1.递归
	2.迭代，每次比较头结点的值
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 1.递归
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	// 2.迭代
	//if list1 == nil {
	//	return list2
	//} else if list2 == nil {
	//	return list1
	//}
	//var p *ListNode
	//if list1.Val < list2.Val {
	//	p = list1
	//	list1 = list1.Next
	//} else {
	//	p = list2
	//	list2 = list2.Next
	//}
	//q := p
	//for list1 != nil || list2 != nil {
	//	if list1 == nil {
	//		p.Next = list2
	//		return q
	//	} else if list2 == nil {
	//		p.Next = list1
	//		return q
	//	}
	//	if list1.Val < list2.Val {
	//		p.Next = list1
	//		p = p.Next
	//		list1 = list1.Next
	//	} else {
	//		p.Next = list2
	//		p = p.Next
	//		list2 = list2.Next
	//	}
	//}
	//return q
}
