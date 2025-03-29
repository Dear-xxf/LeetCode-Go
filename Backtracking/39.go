package Backtracking

/*
   题目描述：
	Middle 组合总和
	给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
	candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
	对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

/*
   解题思路：
	回溯；
	值得注意的点是在写内置函数backtrack的时候，如果将conbination作为参数，是不需要回溯的；
	因为在每一次函数调用都会复制数组，只有当conbination不作为参数使用，不回溯才会导致不同路径结果的堆积；
*/

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	combination := []int{}
	var backtrack func(index, target int)
	backtrack = func(index, target int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		backtrack(index+1, target)
		if target-candidates[index] >= 0 {
			combination = append(combination, candidates[index])
			backtrack(index, target-candidates[index])
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(0, target)
	return res
}
