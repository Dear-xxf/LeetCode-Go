package NormalArray

/*
	题目描述：
	Middle
	最大子数组和
	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。、
	子数组是数组中的一个连续部分。
*/

/*
	解题思路：
	1.动态规划
		用一个数组记录以第i个位置结尾的最大子数组的和，向后动态规划
	2.分治法
*/

func maxSubArray(nums []int) int {
	// 1.动态规划
	res := nums[0]
	fmax := []int{}
	for i := 0; i < len(nums); i++ {
		fmax = append(fmax, nums[i])
		if i == 0 {
			continue
		}
		if fmax[i-1]+nums[i] > fmax[i] {
			fmax[i] = fmax[i-1] + nums[i]
		}
		if fmax[i] > res {
			res = fmax[i]
		}
	}
	return res
}
