package DoublePointer

/*
	题目描述:
	给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
	找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
	返回容器可以储存的最大水量。
	说明：你不能倾斜容器。
*/

/*
	解题思路：
	left指针从左开始，right指针从右开始，那一边height值小，哪一边向中间移动
	直到两个指针相遇
*/

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for right >= left {
		h := min(height[left], height[right])
		curArea := (right - left) * h
		if curArea > maxArea {
			maxArea = curArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
