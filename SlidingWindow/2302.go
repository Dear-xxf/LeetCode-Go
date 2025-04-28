package SlidingWindow

/*
   题目描述：
	Hard 统计得分小于K的子数组数目
	一个数组的 分数 定义为数组之和 乘以 数组的长度。
	比方说，[1, 2, 3, 4, 5] 的分数为 (1 + 2 + 3 + 4 + 5) * 5 = 75 。
	给你一个正整数数组 nums 和一个整数 k ，请你返回 nums 中分数 严格小于 k 的 非空整数子数组数目。
	子数组 是数组中的一个连续元素序列。
*/

/*
   解题思路：
	维护一个滑动窗口 [left, right]，表示当前正在考虑的子数组。
	窗口内维护元素的总和 sum。
	子数组 [left, right] 的得分是：
	score = sum × (right - left + 1)
	score=sum×(right−left+1)
	每次向右扩展 right，累加元素到 sum。
	如果当前得分 score >= k，就不断右移 left，并更新 sum，直到得分小于 k。
	每次窗口合法时，以当前 right 为结尾的合法子数组个数是 right - left + 1，累加到答案中。
*/

func countSubarrays2302(nums []int, k int64) int64 {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	var res int64 = 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for int64(sum)*int64(right-left+1) >= k {
			sum -= nums[left]
			left++
		}
		res += int64(right - left + 1)
	}
	return res
}
