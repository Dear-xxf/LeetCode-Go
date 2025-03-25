package BinaryTree

/*
   题目描述：
	Easy 将有序数组转换为二叉搜索树
	给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
*/

/*
   解题思路：
	递归秒了
*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}
	index := int(len(nums) / 2)
	root := &TreeNode{Val: nums[index]}
	leftNums := nums[:index]
	rightNums := nums[index+1:]
	root.Left = sortedArrayToBST(leftNums)
	root.Right = sortedArrayToBST(rightNums)
	return root
}
