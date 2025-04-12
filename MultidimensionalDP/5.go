package MultidimensionalDP

/*
   题目描述：
	Middle 最长回文子串
	给你一个字符串 s，找到 s 中最长的 回文 子串。
*/

/*
   解题思路：
	一维DP在连续相等的情况下无法判断，例如"ccc"，只能判断为"cc";
	所以使用二维DP，用dp[i][j]表示第i到j位置的子串是否为回文子串。
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > maxLen {
					start = i
					maxLen = j - i + 1
				}
			}
		}
	}
	return s[start : start+maxLen]
}
