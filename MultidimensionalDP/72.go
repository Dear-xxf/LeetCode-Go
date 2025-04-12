package MultidimensionalDP

/*
   题目描述：
	Middle 编辑距离
	给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
	你可以对一个单词进行如下三种操作：
	插入一个字符
	删除一个字符
	替换一个字符
*/

/*
   解题思路：
	MDP题。看似很复杂，拆解到DP状态转移方程却很简单；
	状态定义
		设 dp[i][j] 表示：
		将 word1[0..i-1] 转换为 word2[0..j-1] 所需的最小编辑距离。
	状态转移方程
		如果 word1[i-1] == word2[j-1]：字符一样，不需要操作：
			dp[i][j] = dp[i-1][j-1]
		否则，取三种操作中的最小值 + 1：
			插入：dp[i][j-1] + 1
			删除：dp[i-1][j] + 1
			替换：dp[i-1][j-1] + 1
*/

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			minLen := dp[i-1][j-1]
			if dp[i][j-1] < minLen {
				minLen = dp[i][j-1]
			}
			if dp[i-1][j] < minLen {
				minLen = dp[i-1][j]
			}
			dp[i][j] = minLen + 1
		}
	}
	return dp[len1][len2]
}
