package BinaryTree

/*
   题目描述：
	Middle 二叉搜索树中第K小的元素
	给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
*/

/*
   解题思路：
	同样用中序遍历可以实现。下面采用递归的方式。
	numOfTree记录以当前root结点的树的结点总个数；
	如果左子树的总个数等于k-1，说明当前root结点的val为要找的值
	大于k-1，说明要找的结点在左结点上
	小于k-1，说明要找的结点在右结点上
*/

func kthSmallest(root *TreeNode, k int) int {
	var numOfTree func(root *TreeNode) int
	numOfTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return numOfTree(root.Left) + numOfTree(root.Right) + 1
	}
	if numOfTree(root.Left) == k-1 {
		return root.Val
	} else if numOfTree(root.Left) > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-numOfTree(root.Left)-1)
	}
}
