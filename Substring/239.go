package Substring

/*
	题目描述：
	Hard
	给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
	返回 滑动窗口中的最大值 。
*/

/*
	解题思路：
	毫无思路，第一次做单调队列的题
	单调队列适用于 维护区间内的最值 的问题，尤其是在 滑动窗口、区间动态更新、递增/递减序列处理 相关的题目中
	在本题中，我们使用单调队列来记录最大值所在的索引
	在每一次向右滑动之后，我们都判断新出现的元素是否当前队列中最后一个索引所在元素小，
	如果不是，当前最后一个元素弹出，然后新元素进入队列
	并且在滑动时检查当前队列中第一个元素是否还在窗口中，如果不在了，将其从队列中移除
*/

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	deque := []int{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
