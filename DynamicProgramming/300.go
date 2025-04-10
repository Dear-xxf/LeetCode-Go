package DynamicProgramming

/*
   题目描述：
	Middle 最长递增子序列
	给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/

/*
   解题思路：
	DP
*/

func lengthOfLIS(nums []int) int {
	res := 0
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for index := 0; index < len(nums); index++ {
		dp[index+1] = 1
		for i := 0; i < index; i++ {
			if nums[i] < nums[index] {
				dp[index+1] = max(dp[index+1], dp[i+1]+1)
			}
		}
	}
	for _, value := range dp {
		if res < value {
			res = value
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
