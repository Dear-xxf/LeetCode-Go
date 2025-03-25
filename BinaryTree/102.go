package BinaryTree

/*
   题目描述：
	Middle 二叉树的层序遍历
	给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
*/

/*
   解题思路：
	没有太多思路好写的，利用队列实现
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{[]int{}}
	queue := []*TreeNode{root}
	num := []int{0} // 记录当前结点在第几层
	curLayer := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		nodeLayer := num[0]
		num = num[1:]

		if nodeLayer != curLayer {
			curLayer = nodeLayer
			res = append(res, []int{})
		}
		curLayerList := res[len(res)-1]
		curLayerList = append(curLayerList, node.Val)
		res[len(res)-1] = curLayerList

		if node.Left != nil {
			queue = append(queue, node.Left)
			num = append(num, nodeLayer+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			num = append(num, nodeLayer+1)
		}
	}
	return res
}
