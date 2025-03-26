package BinaryTree

/*
   题目描述：
	Middle 从前序与中序遍历序列构造二叉树
	给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

/*
   解题思路：
	1.递归写法
		递归的想法还是比较简单的，对于preorder的第一个元素就是当前子树的根结点
		对应的，inorder中找到这个值之后，inorder这个值左边的切片是根结点的左子树，右边的切片是根结点的右子树
		分别对根结点的左右子女进行赋值；如此递归下去就可以解决问题
	2.迭代写法
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 1.递归
	if len(preorder) == 0 || preorder[0] == -1 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	rootPosInOrder := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootPosInOrder = i
			break
		}
	}
	leftLen := rootPosInOrder
	rightLen := len(preorder) - rootPosInOrder - 1
	root.Left = buildTree(preorder[1:leftLen+1], inorder[0:leftLen])
	if rightLen != 0 {
		root.Right = buildTree(preorder[leftLen+1:len(preorder)-1], inorder[len(inorder)-rightLen:len(inorder)-1])
	}
	return root
}
