package Stack

import "math"

/*
   题目描述：
	Hard 柱状图中最大的矩形
	给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
	求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

/*
   解题思路：
	思路和接雨水很像，也是给我写恶心了；
	单调栈来解决；栈内存待发现比栈中元素小的元素；
	在遍历过程中，发现当前值比栈顶元素小后，弹出栈顶元素，计算栈顶元素左右可以延申到哪里；
	延申的宽度就是当前index和栈中再前面一个元素位置的距离，因为再上一个元素是原栈顶元素左侧的第一个比其小的元素位置；
	结束后再遍历栈，将剩余栈中位置取出进行计算；最后得到最大值。
*/

func largestRectangleArea(heights []int) int {
	stack := []int{}
	ans := math.MinInt
	for index, value := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > value {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				ans = max(ans, heights[topIndex]*(index-stack[len(stack)-1]))
			} else {
				ans = max(ans, heights[topIndex])
			}
		}
		stack = append(stack, index)
	}
	for index, value := range stack {
		if index != 0 {
			ans = max(ans, heights[value]*(len(heights)-1-stack[index-1]))
		} else {
			ans = max(ans, heights[value]*len(heights))
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
