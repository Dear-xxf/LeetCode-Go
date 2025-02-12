package DoublePointer

/*
	题目描述：
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
	请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

/*
	解题思路：
	两个位置符号left和right，起始位置都是0
	left向左移动，如果遇到元素不是0元素，则与right位置元素相互交换，交换完right++
*/

func moveZeroes(nums []int) {
	right := 0
	for left := 0; left < len(nums); left++ {
		if nums[left] != 0 {
			nums[right] = nums[left]
			if left != right {
				nums[left] = 0
			}
			right++
		}
	}
}
