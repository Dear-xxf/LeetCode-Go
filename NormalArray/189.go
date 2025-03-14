package NormalArray

/*
	题目描述：
	Middle 轮转数组
	给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/

/*
	解题思路：
	非常简单的一道题，对于go语言来说，需要熟悉切片的语法
*/

func rotate(nums []int, k int) {
	k = len(nums) % k
	if len(nums) != 1 {
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	}

	// nums[:k] 从 nums 切片的开始位置到索引 k（不包括 k）之间的所有元素
	// nums[k:] 从索引 k 开始到切片 nums 的结尾之间的所有元素。
}
