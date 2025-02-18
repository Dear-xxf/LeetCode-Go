package Substring

/*
	题目描述：
	给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
	子数组是数组中元素的连续非空序列。
*/

/*
	解题思路：
	请注意！！！双指针在这里是行不通的！因为数组中可能存在负数，所以左指针向左滑动的时候sum可能不是增加而是减少！
	应该使用前缀和＋哈希表的方法
	前缀和定义为sum[i]=nums[0]+nums[1]+⋯+nums[i]
	那么，对于任意一个子数组 nums[i...j]，它的和就可以表示为sum[j]−sum[i−1]
	这就意味着如果我们知道当前的前缀和 sum[j]，那么只需要看之前是否存在某个前缀和 sum[i-1] 等于 sum[j] - k
	使用哈希表来存储前缀和
*/

func SubarraySum(nums []int, k int) int {
	currentSum := 0
	count := 0
	prefixSumMap := map[int]int{0: 1} // 初始化：前缀和0出现1次

	for _, num := range nums {
		currentSum += num
		if freq, exists := prefixSumMap[currentSum-k]; exists {
			count += freq
		}
		prefixSumMap[currentSum]++
	}

	return count
}
