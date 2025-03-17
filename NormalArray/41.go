package NormalArray

/*
	题目描述：
	Hard 缺失的第一个正数
	给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
	请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案
*/

/*
	解题思路：
	原地哈希
	这道题目的关键在于发现最大正整数的值不可能超过数组的len
	于是将数组中每一个正整数的值放在其对应的位置上，比如将3放在索引2位置上
	这样最后通过一个循环来判断第一个不匹配的索引即可
*/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 1. 置换到正确的位置
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1] // 交换到正确位置
		}
	}
	// 2. 查找第一个不匹配的位置
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
