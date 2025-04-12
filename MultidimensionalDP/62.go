package MultidimensionalDP

/*
   题目描述：
	Middle 不同路径
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
	问总共有多少条不同的路径？
*/

/*
   解题思路：
	很简单的一道题，到达dp[i][j]位置的路径为dp[i-1][j]和dp[i][j-1]的和。
*/

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row >= 1 {
				dp[row][col] += dp[row-1][col]
			}
			if col >= 1 {
				dp[row][col] += dp[row][col-1]
			}
		}
	}
	return dp[m-1][n-1]
}
