package BinarySearch

/*
   题目描述：
	Easy 搜索插入位置
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
	请必须使用时间复杂度为 O(log n) 的算法。
*/

/*
   解题思路：
	套用二分法即可，即不断用二分法逼近查找第一个大于等于 target 的下标；
	语法上值得关注的是除以2寻找mid的操作使用位运算来执行。
*/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left // 除以2使用位运算
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
