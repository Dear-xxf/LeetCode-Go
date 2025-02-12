package leetcode

/*
	题目描述：
	给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。
	一次操作中，你可以删除 nums 中的最小元素。
	你需要使数组中的所有元素都大于或等于 k ，请你返回需要的 最少 操作次数。
*/

/*
	解题思路：
	遍历数组，当前元素的值比k小的时候，num++，找到所有比k小的数字即可
*/

func minOperations(nums []int, k int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			num++
		}
	}
	return num
}
