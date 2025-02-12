package leetcode

import "container/heap"

/*
	题目描述：
	给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。
	一次操作中，你将执行：
	选择 nums 中最小的两个整数 x 和 y 。
	将 x 和 y 从 nums 中删除。
	将 min(x, y) * 2 + max(x, y) 添加到数组中的任意位置。
	注意，只有当 nums 至少包含两个元素时，你才可以执行以上操作。
	你需要使数组中的所有元素都大于或等于 k ，请你返回需要的 最少 操作次数。
*/

/*
	解题思路：
	优先处理最小的两个数
	使用 最小堆（min-heap）来维护数组中的元素，能够高效获取和删除最小的两个数。
	在 Go 中，使用 container/heap 包来实现最小堆。

	在这题中，熟悉了解heap包的使用比较重要
*/

func minOperations2(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	operations := 0

	for h.Len() > 1 && (*h)[0] < k {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)

		newNum := x*2 + y
		heap.Push(h, newNum)
		operations++
	}

	if (*h)[0] < k {
		return -1
	}
	return operations
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
