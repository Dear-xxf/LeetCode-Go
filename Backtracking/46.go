package Backtracking

/*
   题目描述：
	Middle 全排列
	给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

/*
   解题思路：
	刷力扣第一次做回溯的题目；有一点类似于递归
	1.递归地选择未使用的数字，并加入当前排列。
	2.当排列长度等于 nums 长度时，记录结果。
	3.回溯撤销选择，继续尝试其他可能性
*/

func permute(nums []int) [][]int {
	res := [][]int{}
	var backTrack func(path []int, used []bool)
	backTrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, num)
			backTrack(path, used)
			// 撤销这一步很重要
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack([]int{}, make([]bool, len(nums)))
	return res
}
