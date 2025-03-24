package BinaryTree

/*
   题目描述：
	Easy 二叉树的中序遍历
	给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/

/*
   解题思路：
	通过栈来记录每一个子树的根节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	// 建立一个栈用于存储每一个子树的根节点
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) != 0 {
		// 如果存在左子树，继续向下遍历
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curNode.Val)
		curNode = curNode.Right
	}
	return res
}
