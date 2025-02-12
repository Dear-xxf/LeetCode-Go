package leetcode

import "math"

/*
	题目描述：
	给你一个 非负 整数数组 nums 和一个整数 k 。
	如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。
	请你返回 nums 中 最短特别非空子数组的长度，如果特别子数组不存在，那么返回 -1 。
*/

/*
	解题思路：
	通过双指针定义一个动态窗口，用变量 currentOr 表示窗口内元素的按位或值。
	当窗口的 currentOr ≥ k 时，记录当前窗口长度并尝试缩小窗口以寻找更短的子数组。
	如果遍历后仍未找到满足条件的子数组，则返回 -1；
	否则返回最短长度。这种方法充分利用按位或的单调性，高效解决问题，时间复杂度为 O(n)。
*/

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	minLength := math.MaxInt32

	for start := 0; start < n; start++ {
		currentOR := 0
		for end := start; end < n; end++ {
			currentOR |= nums[end]
			if currentOR >= k {
				minLength = min(minLength, end-start+1)
				break
			}
		}
	}

	if minLength == math.MaxInt32 {
		return -1 // 如果没有找到特别子数组
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
