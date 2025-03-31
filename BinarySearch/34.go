package BinarySearch

/*
   题目描述：
	Middle 在排序数组中查找元素的第一个和最后一个位置
	给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回 [-1, -1]。
	你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

/*
   解题思路：
	实现一个内置函数，搜索的是target在nums中出现的下标最小的位置，如果不存在，就返回应该插入的位置；
	（即sort.SearchInts()函数功能）
	左下标可由此找到，右下标通过SearchInts(nums, target + 1) - 1来找到
	注意searchInt函数可以找到target在nums中下标最小的位置；通过不在等于时返回做到，在等于的时候依然向左查询，直到找到最左边的位置。
*/

func searchRange(nums []int, target int) []int {
	left := searchInt(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := searchInt(nums, target+1) - 1
	return []int{left, right}
}

func searchInt(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
