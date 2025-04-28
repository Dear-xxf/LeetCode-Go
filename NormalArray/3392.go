package NormalArray

/*
   题目描述：
	Easy 统计符合条件长度为3的子数组数目
	给你一个整数数组 nums ，请你返回长度为 3 的 子数组 的数量，满足第一个数和第三个数的和恰好为第二个数的一半。
	子数组 指的是一个数组中连续 非空 的元素序列。
*/

/*
   解题思路：
	遍历一次即可。
*/

func countSubarrays(nums []int) int {
	res := 0
	for index := 0; index < len(nums)-2; index++ {
		if 2*nums[index]+2*nums[index+2] == nums[index+1] {
			res++
		}
	}
	return res
}
