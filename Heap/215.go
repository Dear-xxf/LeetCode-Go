package Heap

/*
   题目描述：
	Middle 数组中的第K个最大元素
	给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
	请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
	你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

/*
   解题思路：
	堆排序
*/

func findKthLargest(nums []int, k int) int {
	constructHeap(nums, len(nums))
	for i := 0; i < k; i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		handleNode(nums, 0, len(nums)-i-1)
	}
	return nums[0]
}

func constructHeap(nums []int, heapsize int) {
	for i := heapsize/2 - 1; i > -1; i-- {
		handleNode(nums, i, heapsize)
	}
}

func handleNode(nums []int, i, heapsize int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < heapsize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapsize && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		handleNode(nums, largest, heapsize)
	}
}
