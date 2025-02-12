package DoublePointer

import "sort"

/*
	题目描述：
	给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
*/

/*
	解题思路：
	在三重循环的背景下，考虑到后面两个指针j和k
	当j在右移时，整体的和变大；当k在左移时，整体的和变小
	所以，将后面的两层循环控制为一层，当和超过0的时候，k左移，当和小于0的时候，j右移
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		for j := i; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if k > j && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}

		}
	}
	return ans
}
