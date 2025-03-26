package BinaryTree

/*
   题目描述：
	Middle 二叉树的右视图
	给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

/*
   解题思路：
	DFS 优先遍历右子女
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	stack := []*TreeNode{root}
	depth := []int{0}
	curDepth := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nodeDepth := depth[len(depth)-1]
		depth = depth[:len(depth)-1]
		if nodeDepth > curDepth {
			res = append(res, node.Val)
			curDepth = nodeDepth
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			depth = append(depth, nodeDepth+1)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			depth = append(depth, nodeDepth+1)
		}
	}
	return res
}
