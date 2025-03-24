package BinaryTree

/*
   题目描述：
	Easy 对称二叉树
	给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

/*
   解题思路：
	做的时候竟然没什么思路，一开始觉得同样使用DFS遍历即可，但是并不满足题目的要求
	使用递归函数解决，建立两个指针指向根结点的左右子女
	递归比较左子女的左子树和右子女的右子树，以及左子女的右子树和右子女的左子树
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
