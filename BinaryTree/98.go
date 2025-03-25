package BinaryTree

import "math"

/*
   题目描述：
	Middle 验证二叉搜索树
	给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
	有效 二叉搜索树定义如下：
	节点的左子树只包含 小于 当前节点的数。
	节点的右子树只包含 大于 当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
*/

/*
   解题思路：
	中序遍历，每一次从栈中取出来的值和上一个值进行比较，如果比上一个值小，则不满足二叉搜索树的条件
	同时每一次从栈中时取元素时，更新num值用于比较
*/

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	curNode := root
	num := math.MinInt
	for len(stack) > 0 || curNode != nil {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curNode.Val <= num {
			return false
		}
		num = curNode.Val

		curNode = curNode.Right
	}
	return true
}
