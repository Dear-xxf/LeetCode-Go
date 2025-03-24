package BinaryTree

/*
   题目描述：
	Easy 二叉树的最大深度
	给定一个二叉树 root ，返回其最大深度。
	二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
*/

/*
   解题思路：
	DFS
	很久没有写DFS，这次再写感觉非常生疏，只记得用栈来实现了
	借这个题好好熟悉一下
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	depthStack := []int{1} // 记录每个节点对应的深度
	maxDepth := 0
	for len(stack) > 0 {
		// 取出当前节点
		node := stack[len(stack)-1]
		depth := depthStack[len(depthStack)-1] // 取出当前节点的深度
		// 出栈
		stack = stack[:len(stack)-1]
		depthStack = depthStack[:len(depthStack)-1]

		if depth > maxDepth {
			maxDepth = depth
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			depthStack = append(depthStack, depth+1)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			depthStack = append(depthStack, depth+1)
		}
	}

	return maxDepth
}
