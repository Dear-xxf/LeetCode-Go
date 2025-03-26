package BinaryTree

/*
   题目描述：
	Middle 路径总和Ⅲ
	给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
	路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

/*
   解题思路：
	前缀和的思路是利用哈希表 prefixSumCount 记录从根节点到当前节点的路径和出现的次数，然后在遍历二叉树时检查是否存在某个祖先节点，使得从该祖先节点到当前节点的路径和等于 targetSum。
	核心公式：当前前缀和 − 目标前缀和 = targetSum
	当前前缀和−目标前缀和=targetSum
	如果 prefixSumCount[curSum - targetSum] 存在，就说明有多少条路径符合要求。
	遍历过程中维护哈希表、递归搜索左右子树、回溯恢复状态，最终高效计算路径总数，避免 O(n²) 的暴力解法，使时间复杂度优化到 O(n)。
	（思路由chatGPT生成，自己实现时除了暴力没有太好的思路）
*/

func pathSum(root *TreeNode, targetSum int) int {
	prefixSumCount := map[int]int{0: 1} // 记录前缀和出现的次数

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, curSum int) int {
		if node == nil {
			return 0
		}
		// 更新当前路径的前缀和
		curSum += node.Val
		// 计算从某个祖先节点到当前节点的路径是否满足 targetSum
		count := prefixSumCount[curSum-targetSum]
		// 记录当前前缀和
		prefixSumCount[curSum]++
		// 递归计算左右子树的路径数
		count += dfs(node.Left, curSum)
		count += dfs(node.Right, curSum)
		// 回溯：移除当前前缀和，避免影响其他路径
		prefixSumCount[curSum]--
		return count
	}
	return dfs(root, 0)
}
