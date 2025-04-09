package DynamicProgramming

import "math"

/*
   题目描述：
	Middle 零钱兑换
	给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
	计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
	你可以认为每种硬币的数量是无限的。
*/

/*
   解题思路：
	DP;
	1. 定义状态
		dp[i] 表示凑出金额 i 所需的最小硬币数。
	2. 初始化
		dp[0] = 0，表示金额为 0 时需要 0 枚硬币。
		其他 dp[i] 初始化为amount + 1，表示“还没计算出来”或“不可能”。
	3. 状态转移方程
		对于每个金额 i，我们尝试所有硬币 coin：
			dp[i] = min(dp[i], dp[i - coin] + 1)  // 如果 i - coin >= 0
		意思是：
			如果我想凑出金额 i，并且我用了一个面值为 coin 的硬币，
			那么我只需要再加上凑出金额 i - coin 所需的最小硬币数，再加 1。
*/

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt
		for _, coin := range coins {
			if i == coin {
				dp[i] = 1
				continue
			}
			if i-coin > 0 && dp[i-coin] != -1 {
				if dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
