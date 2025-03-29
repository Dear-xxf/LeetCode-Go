package Backtracking

/*
   题目描述：
	Middle 括号生成
	数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

/*
   解题思路：
	递归构建字符串：逐步尝试添加 ( 或 )，直到字符串长度达到 2n。
	约束条件（剪枝）：
		左括号：只要已使用的左括号 < n，就可以继续加 (。
		右括号：只有已使用的右括号 < 左括号 时，才能加 )（否则会形成无效组合，如 ())）。
	终止条件：当字符串长度达到 2n 时，加入结果列表。
*/

func generateParenthesis(n int) []string {
	res := []string{}
	combination := ""
	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		if left < n-right { // 已经使用过的左括号数量小于已经使用过的右括号数量
			return
		}
		if left > n || right < 0 {
			return
		}
		if len(combination) == 2*n {
			if left == n {
				res = append(res, combination)
			}
			return
		}
		combination += "("
		backtrack(left+1, right)
		combination = combination[:len(combination)-1]

		combination += ")"
		backtrack(left, right-1)
		combination = combination[:len(combination)-1]
	}
	backtrack(0, n)
	return res
}
