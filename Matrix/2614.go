package Matrix

/*
	题目描述：
	Easy 对角线上的质数
	给你一个下标从 0 开始的二维整数数组 nums 。
	返回位于 nums 至少一条 对角线 上的最大 质数 。如果任一对角线上均不存在质数，返回 0 。
*/

/*
	解题思路：
	简单题，没有太复杂的要考虑，分别考虑主对角线和副对角线即可
*/

func diagonalPrime(nums [][]int) int {
	res := 0
	// 主对角线
	i := 0
	j := 0
	for i < len(nums) && j < len(nums[0]) {
		if isPrime(nums[i][j]) && nums[i][j] > res {
			res = nums[i][j]
		}
		i++
		j++
	}
	// 副对角线
	i = 0
	j = len(nums[0]) - 1
	for i < len(nums) && j > -1 {
		if isPrime(nums[i][j]) && nums[i][j] > res {
			res = nums[i][j]
		}
		i++
		j--
	}
	return res
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	// 最后要判断数字是否大于2
	return num >= 2
}
