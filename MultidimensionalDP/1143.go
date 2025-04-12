package MultidimensionalDP

/*
   题目描述：
	Middle 最长公共子序列
	给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
	一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
	例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
	两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/

/*
   解题思路：
	构造的过程很难想；
	设 dp[i][j] 表示：
	text1[0..i-1] 与 text2[0..j-1] 的最长公共子序列长度
	注意这里的 i 和 j 是从 1 开始对应字符串的前缀长度（即第 i 个字符在字符串中是 i-1 下标）
	状态转移方程
		如果 text1[i-1] == text2[j-1]：
		说明这个字符可以纳入公共子序列
		dp[i][j] = dp[i-1][j-1] + 1
	否则：我们可以跳过 text1[i-1] 或 text2[j-1]
	dp[i][j] = max(dp[i-1][j], dp[i][j-1])
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
