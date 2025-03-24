package BinaryTree

/*
   题目描述：
	Easy 翻转二叉树
	给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
*/

/*
   解题思路：
	DFS遍历，每一次从栈顶弹出元素的时候交换该元素的左右子女
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		temp := curNode.Left
		curNode.Left = curNode.Right
		curNode.Right = temp

		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	return root
}
