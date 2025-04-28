package SlidingWindow

import "math"

/*
   题目描述：
	Middle 统计最大元素出现至少K次的子数组
	给你一个整数数组 nums 和一个 正整数 k 。
	请你统计有多少满足 「 nums 中的 最大 元素」至少出现 k 次的子数组，并返回满足这一条件的子数组的数目。
	子数组是数组中的一个连续元素序列。
*/

/*
   解题思路：
	很简单的滑动窗口题；右指针向右滑动，当最大值总数到达k个之后，左指针向左滑动。
*/

func countSubarrays(nums []int, k int) int64 {
	// 找到nums数组中的最大值
	maxElement := math.MinInt
	for _, value := range nums {
		if value > maxElement {
			maxElement = value
		}
	}

	left := 0
	var res int64 = 0
	count := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == maxElement {
			count++
		}
		for count == k {
			res += int64(len(nums) - right)
			if nums[left] == maxElement {
				count--
			}
			left++
		}
	}
	return res
}
