package BinaryTree

/*
   题目描述：
	Middle 二叉树展开为链表
	给你二叉树的根结点 root ，请你将它展开为一个单链表：
	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

/*
   解题思路：
	前序遍历可以解决，还是不够熟悉三种遍历在go语言中的写法。
	这里对从根结点开始每一个结点，找到其右子女在链表中的正确位置，也就是左子树中最右边的结点
*/

func flatten(root *TreeNode) {
	node := root
	for node != nil {
		if node.Left != nil {
			next := node.Left
			previous := next
			for previous.Right != nil {
				previous = previous.Right
			}
			previous.Right = node.Right
			node.Left, node.Right = nil, next
		}
		node = node.Right
	}
}
