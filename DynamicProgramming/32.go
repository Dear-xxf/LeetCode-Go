package DynamicProgramming

/*
   题目描述：
	Hard 最长有效括号
	给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
*/

/*
   解题思路：
	定义：
		dp[i] 表示以第 i 个字符结尾的最长有效括号长度。
	转移规则：
		如果 s[i] == ')'：
		若 s[i-1] == '('，那么 dp[i] = dp[i-2] + 2
		若 s[i-1] == ')' 并且 s[i - dp[i-1] - 1] == '('，则 dp[i] = dp[i-1] + dp[i - dp[i-1] - 2] + 2
	代码中需要注意边界条件，不能溢出，即i == 2和previousIndex == 0情况。
*/

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	maxLen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				previousIndex := i - 1 - dp[i-1]
				if previousIndex >= 0 && s[previousIndex] == '(' {
					if previousIndex >= 1 {
						dp[i] = dp[i-1] + 2 + dp[previousIndex-1]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
