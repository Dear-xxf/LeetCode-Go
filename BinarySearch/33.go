package BinarySearch

/*
   题目描述：
	Middle 搜索旋转排序数组
	整数数组 nums 按升序排列，数组中的值 互不相同 。
	在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
	给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
	你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
*/

/*
   解题思路：
	二分查找的基本思想：
		由于数组是部分有序的，我们可以利用二分查找来减少搜索范围。
		每次找到中点 mid，并判断左半部分或右半部分是否是有序的。
	确定有序区间：
		如果 nums[left] <= nums[mid]，说明左半部分是有序的。
		如果 nums[mid] <= nums[right]，说明右半部分是有序的。
	判断 target 是否在有序区间内：
		如果 target 在有序区间内，则移动 left 或 right 继续搜索。
		否则，在另一半区间继续搜索
*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		// 找到目标值，返回索引
		if nums[mid] == target {
			return mid
		}

		// 判断哪一半是有序的
		if nums[left] <= nums[mid] { // 左半部分有序
			if nums[left] <= target && target < nums[mid] { // target 在左半部分
				right = mid - 1
			} else { // target 在右半部分
				left = mid + 1
			}
		} else { // 右半部分有序
			if nums[mid] < target && target <= nums[right] { // target 在右半部分
				left = mid + 1
			} else { // target 在左半部分
				right = mid - 1
			}
		}
	}
	return -1 // 没找到
}
