package DynamicProgramming

/*
   题目描述：
	Middle 乘积最大子数组
	给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
	测试用例的答案是一个 32-位 整数。
*/

/*
   解题思路：
	DP;这题不同于其他DP题目在于需要同时考虑最大值和最小值；
	原因在于nums数组中可能出现负数，负数翻转后可能成为最大的正数。
*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal, minVal, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal
		}
		maxVal = max(nums[i], maxVal*nums[i])
		minVal = min(nums[i], minVal*nums[i])
		res = max(res, maxVal)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
