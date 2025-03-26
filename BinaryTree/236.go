package BinaryTree

/*
   题目描述：
	Middle 二叉树的最近公共祖先
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
	百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

/*
   解题思路：
	1.递归，不做过多解释
	2.记录父结点；
		遍历二叉树，记录每一个结点的父结点信息到哈希表中
		之后再用另一个哈希表记录p结点的所有祖先信息，在遍历q结点的时候判断祖先有没有重合即可
	第二种方法实际的内存开销更大且耗时更多，表现反而不如递归
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 1.递归
	//if root == nil {
	//	return nil
	//}
	//if root.Val == p.Val || root.Val == q.Val {
	//	return root
	//}
	//left := lowestCommonAncestor(root.Left, p, q)
	//right := lowestCommonAncestor(root.Right, p, q)
	//if left != nil && right != nil {
	//	return root
	//}
	//if left == nil {
	//	return right
	//}
	//return left

	// 2.记录父结点
	if root == nil {
		return nil
	}
	hash := map[*TreeNode]*TreeNode{root: nil}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
			hash[node.Left] = node
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			hash[node.Right] = node
		}
	}
	parent_of_p := map[*TreeNode]bool{}
	for p != nil {
		parent_of_p[p] = true
		p = hash[p]
	}
	for q != nil {
		if parent_of_p[q] == true {
			return q
		}
		q = hash[q]
	}
	return nil
}
