package DynamicProgramming

/*
   题目描述：
	Middle 单词拆分
	给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
	注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
*/

/*
   解题思路：
	DP;
	第一个循环遍历字符串s的每一个位置；
	第二个循环对每一个word进行遍历，看i-len(word)位置是否为true，若为true，dp[i]也为true。
*/

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			l := len(word)
			if i >= l && s[i-l:i] == word && dp[i-l] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
