package BinaryTree

import "math"

/*
   题目描述：
	Hard 二叉树中的最大路径和
	二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
	路径和 是路径中各节点值的总和。
	给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

/*
   解题思路：
	难得非常简单的Hard题，递归秒了
	递归内容是从每一个结点出发，做两件事；
	第一件用左右子树递归上来的值计算经过该结点的路径最大值并更新
	第二件向上递归返回该结点作为子树时向父结点能提供的最大值
*/

func maxPathSum(root *TreeNode) int {
	var maxSupport func(root *TreeNode) int
	maxSum := math.MinInt
	maxSupport = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSupport := max(maxSupport(root.Left), 0)
		rightSupport := max(maxSupport(root.Right), 0)

		pathSum := root.Val + leftSupport + rightSupport
		maxSum = max(maxSum, pathSum)

		return root.Val + max(leftSupport, rightSupport)
	}
	maxSupport(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
