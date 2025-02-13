package DoublePointer

/*
	题目描述：
	hard
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

/*
	解题思路：
	方法一：双指针
		这道题目的核心在于发现当前位置上能接住的雨水与什么有关
		可以认定其与左右两边最大值相关
		所以在实现的时候使用双指针，一个从左一个从右
		当左指针的值小于右指针的值时，说明左边的最大值一定小于右边的最大值，左边的指针向右移动
		当右指针的值小于等于左指针的值时，说明右边的最大值一定等于小于左边的，右边的指针向左移动
*/

func trap(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	leftMax := height[left]
	rightMax := height[right]
	for right > left {
		// 更新leftMax值和rightMax值
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}
		// 如果左指针的值小于右指针的值
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		}
		// 如果右指针的值小于左指针的值
		if height[left] >= height[right] {
			ans += rightMax - height[right]
			right--
		}
		// 两个指针相遇的时候，退出循环
		if left == right {
			break
		}
	}
	return ans
}
