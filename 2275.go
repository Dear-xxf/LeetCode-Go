package leetcode

/*
	题目描述：
	对数组 nums 执行 按位与 相当于对数组 nums 中的所有整数执行 按位与 。
	例如，对 nums = [1, 5, 3] 来说，按位与等于 1 & 5 & 3 = 1 。
	同样，对 nums = [7] 而言，按位与等于 7 。
	给你一个正整数数组 candidates 。计算 candidates 中的数字每种组合下 按位与 的结果。
	返回按位与结果大于 0 的 最长 组合的长度。
*/

/*
	解题思路：
	1.每个数字可以表示为一个二进制数。对于每一位，按位与的结果是否大于 0 取决于该位是否有至少一个 1。
	2.遍历数组 candidates 中的所有数字。对于每个数字，检查其二进制表示的每一位，并统计所有数字中每个位上 1 出现的次数。
	3.对每个位上统计的 1 的数量，取其中的最大值。这就是按位与结果大于 0 的最长组合长度。
*/

func LargestCombination(candidates []int) int {
	maxBitCount := 0
	// 循环32次 因为二进制int32位 对每一位上1的数量进行统计
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range candidates {
			if num&(1<<i) != 0 {
				count++
			}
		}
		if count > maxBitCount {
			maxBitCount = count
		}
	}
	return maxBitCount
}
