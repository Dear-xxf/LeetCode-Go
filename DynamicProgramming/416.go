package DynamicProgramming

/*
   题目描述：
	Middle 分割等和子集
	给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/

/*
   解题思路：
	这道题目实现的时候想到了找到和为sum/2的子数组，但是在构建状态转移方程的时候没有处理好细节。
	下面思路与解法是gpt生成。
	1. 问题转化
		如果总和为 sum，我们要判断能否从数组中选出一些数，使它们的和为 sum / 2。
		因为如果能做到，那么剩下的那些数之和也一定是 sum / 2，就可以分成两个相等的子集。
		所以目标变为：在 nums 中选出若干个数（每个数只能用一次），使它们的和为 sum / 2。
	2. 先判断是否可行
		如果 sum 是奇数，直接返回 false。
		否则设目标和 target = sum / 2。
	3. 动态规划解法（0-1 背包）
		使用一个一维布尔数组 dp，其中 dp[i] 表示是否可以选出一些数，使得它们的和为 i。
		初始时 dp[0] = true（和为 0 总是可以的）。
		然后遍历 nums 中的每个数，对于每个数，从 target 到 num 逆序遍历：
		dp[j] = dp[j] || dp[j - num]
		表示：如果之前能组成 j - num，那么现在就能组成 j。
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		// 从目标值 target 递减到当前数字 num
		// 这样能保证每个数字只被使用一次（避免重复使用 num）
		for j := target; j >= num; j-- {
			// 如果 dp[j - num] 是 true，说明存在某些数字之和为 j - num，
			// 加上当前 num 就可以组成 j，所以 dp[j] 应该也为 true。
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
