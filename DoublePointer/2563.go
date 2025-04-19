package DoublePointer

import "sort"

/*
   题目描述：
	Middle 统计公平数对的数目
	给你一个下标从 0 开始、长度为 n 的整数数组 nums ，和两个整数 lower 和 upper ，返回 公平数对的数目 。
	如果 (i, j) 数对满足以下情况，则认为它是一个 公平数对 ：
	0 <= i < j < n，且
	lower <= nums[i] + nums[j] <= upper
*/

/*
   解题思路：
	排序数组：
		首先对 nums 进行排序，使得之后可以利用二分查找快速定位满足条件的区间。
	遍历每个元素 i：
		对于当前元素 nums[i]，找所有满足 lower - nums[i] <= nums[j] <= upper - nums[i] 且 j > i 的 j。
	二分查找 + 前缀计数：
		使用 sort.Search 找到满足上述条件的 j 的区间：
		left 是第一个满足 nums[j] >= lower - nums[i] 的下标
		right 是第一个满足 nums[j] > upper - nums[i] 的下标
		所以合法的 j 范围是 [left, right)，其中还需要确保 j > i，即从 max(left, i+1) 开始统计。
	累加结果：
		对每个 i 统计合法的 j 的个数，然后将它们加总起来就是答案。
*/

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var res int64 = 0
	n := len(nums)
	for i := 0; i < n; i++ {
		// nums[j] ∈ [lower - nums[i], upper - nums[i]], 且 j > i
		left := sort.Search(n, func(j int) bool {
			return nums[j] >= lower-nums[i]
		})
		right := sort.Search(n, func(j int) bool {
			return nums[j] > upper-nums[i]
		})
		// 剔除掉 j <= i 的部分
		cnt := right - left
		if left <= i && i < right {
			cnt-- // 自己不能和自己配对
		}
		res += int64(cnt)
	}
	// 每对 (i, j) 被统计两次，除以2
	return res / 2
}
