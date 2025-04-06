package BinarySearch

/*
   题目描述：
	Hard 寻找两个正序数组的中位数
	给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
	算法的时间复杂度应该为 O(log (m+n)) 。
*/

/*
   解题思路：
	第一个想到的是合并数组后排序，但是时间复杂度为O(m+n)
	解题核心思想：找第 k 小的元素（第 (m+n)/2 个）
		中位数是第 k = (m + n)/2 个或中间两个数的平均值（如果总数是偶数）。
		用一个辅助函数 getKthElement(nums1, nums2, k) 来找第 k 小的数。
	getKthElement 的思路
		每次从两个数组中各取 k/2 个元素出来比较。
		谁的第 k/2 个元素小，就说明它前面这部分都不是第 k 小的，全部可以排除掉。
		然后在剩下的数组中继续找第 k - 去掉的元素数 小的数。
		这个过程不断缩小问题规模，类似二分查找，时间复杂度是 O(log(k))。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 { // 奇数情况下
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else { // 偶数情况下
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	half1 := min(k/2, len(nums1)) - 1
	half2 := min(k/2, len(nums2)) - 1
	pivot1 := nums1[half1]
	pivot2 := nums2[half2]
	// 核心逻辑，如果num1数组第k/2−1个值小于nums2数组中的，则说明nums1数组前k/2−1个数字都不会是第k个值
	if pivot1 <= pivot2 {
		return getKthElement(nums1[half1+1:], nums2, k-(half1+1))
	} else { // 如果nums2第k/2−1个值小于nums1数组中的，处理nums2
		return getKthElement(nums1, nums2[half2+1:], k-(half2+1))
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
