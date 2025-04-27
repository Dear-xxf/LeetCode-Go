package SlidingWindow

/*
   题目描述：
	Middle 统计完全子数组的数目
	给你一个由 正 整数组成的数组 nums 。
	如果数组中的某个子数组满足下述条件，则称之为 完全子数组 ：
	子数组中 不同 元素的数目等于整个数组不同元素的数目。
	返回数组中 完全子数组 的数目。
	子数组 是数组中的一个连续非空序列。
*/

/*
   解题思路：
	用双指针 left 和 right 维护一个滑动窗口 [left, right]。
	用哈希表实时记录窗口内不同元素的数量。
	只要窗口内包含了所有种类（当前窗口的不同元素种类 == 总种类），那么以当前 right 结尾的子数组 [i, right]，i 从 left 到 right，都是合法的。
*/

func countCompleteSubarrays(nums []int) int {
	// 先统计nums中一共有几种数字
	totalKinds := 0
	freq := map[int]int{}
	for _, num := range nums {
		if freq[num] == 0 {
			totalKinds++
		}
		freq[num]++
	}

	// 用于存储窗口内不同元素的个数
	windowCount := make(map[int]int)
	left := 0
	result := 0

	// 通过右指针扩展窗口
	for right := 0; right < len(nums); right++ {
		// 扩展窗口，增加当前元素的计数
		windowCount[nums[right]]++

		// 当窗口内的不同元素种类等于总种类时，说明满足条件
		for len(windowCount) == totalKinds {
			// 计算满足条件的子数组数量
			result += len(nums) - right

			// 收缩窗口，缩小左边界
			windowCount[nums[left]]--
			if windowCount[nums[left]] == 0 {
				delete(windowCount, nums[left])
			}
			left++
		}
	}

	return result
}
