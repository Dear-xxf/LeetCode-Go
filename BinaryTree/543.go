package BinaryTree

import "math"

/*
   题目描述：
	Easy 二叉树的直径
	给你一棵二叉树的根节点，返回该树的 直径 。
	二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
	两节点之间路径的 长度 由它们之间边数表示。
*/

/*
   解题思路：
	思路还是围绕递归，把每一条路径理解成一个结点为根结点，从其左儿子和右儿子向下遍历的路径拼接得到
	语法上进一步熟悉一下go语言局部函数的写法
	总结感觉是披着Easy皮的Middle题
*/

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var lenFromBottom func(root *TreeNode) int // 先声明
	lenFromBottom = func(root *TreeNode) int { // 再赋值
		if root == nil {
			return 0
		}
		L := lenFromBottom(root.Left)
		R := lenFromBottom(root.Right)
		if L+R+1 > res {
			res = L + R + 1
		}
		return int(math.Max(float64(L), float64(R))) + 1
	}
	lenFromBottom(root)
	return res - 1
}
