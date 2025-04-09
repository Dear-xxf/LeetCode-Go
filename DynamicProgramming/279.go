package DynamicProgramming

/*
   题目描述：
	Middle 完全平方数
	给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
	完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/

/*
   解题思路：
	DP；
	对于每个 i 从 1 到 n，尝试所有可能的平方数 j*j <= i：
	dp[i] = min(dp[i], dp[i - j*j] + 1)
	如果使用了一个 j*j，那么剩下的 i - j*j 用最优解来填，加上当前这个平方数就是一个新的可能方案。
*/

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = i
		for s := 1; s*s <= i; s++ {
			if dp[i] > dp[i-s*s]+1 {
				dp[i] = dp[i-s*s] + 1
			}
		}
	}
	return dp[n]
}
