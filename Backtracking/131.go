package Backtracking

/*
   题目描述：
	Middle 分割回文串
	给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
*/

/*
   解题思路：
	一次写完的回溯算法；官方题解给出两种算法
	1.回溯 + 动态规划
	2.回溯 + 记忆化搜索
	我实现的类似第二种方法，但是对是否回文的结果进行存储，而是每一次都放在backtrack里执行；
	优化的点在于加上记忆化搜索或者加上动态规划预处理
*/

func partition(s string) [][]string {
	res := [][]string{}
	substrings := []string{}
	var backtrack func(indexIns, indexIntemp int)
	// 一个代表在原字符串s中第几个字符，一个代表在暂存结果中第几个子串
	backtrack = func(indexIns, indexIntemp int) {
		if indexIns == len(s) {
			if indexIntemp == len(substrings) {
				copySubstrings := make([]string, len(substrings))
				copy(copySubstrings, substrings)
				res = append(res, copySubstrings)
			}
			return
		}
		if len(substrings) == indexIntemp { // 说明当前子串为新分割出的子串
			// 用当前字符赋值新的子串
			substrings = append(substrings, string(s[indexIns]))
			backtrack(indexIns+1, indexIntemp)
			backtrack(indexIns+1, indexIntemp+1)
			substrings = substrings[:len(substrings)-1]
		} else {
			temp := substrings[len(substrings)-1]
			substrings[len(substrings)-1] += string(s[indexIns])
			backtrack(indexIns+1, indexIntemp)
			if checkPalindrome(substrings[len(substrings)-1]) {
				backtrack(indexIns+1, indexIntemp+1)
			}
			substrings[len(substrings)-1] = temp
		}
	}
	backtrack(0, 0)
	return res
}

func checkPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
