package DynamicProgramming

/*
   题目描述：
	Middle 打家劫舍
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
	如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
*/

/*
   解题思路：
	注意到每一个位置i的值都是由i-2和i-3两个位置分别加上当前nums[i]的值得来；选择这两个值中的最大值作为dp[i]。
*/

func rob(nums []int) int {
	if len(nums) < 3 {
		res := nums[0]
		for i := 0; i < len(nums); i++ {
			if nums[i] > res {
				res = nums[i]
			}
		}
		return res
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]
	res := dp[1]
	if dp[2] > dp[1] {
		res = dp[2]
	}
	for i := 3; i < len(nums); i++ {
		if dp[i-3] > dp[i-2] {
			dp[i] = dp[i-3] + nums[i]
		} else {
			dp[i] = dp[i-2] + nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
