package Backtracking

/*
   题目描述：
	Middle 子集
	给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
	解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

/*
   解题思路：
	回溯；关键在于每个子集的构成，在于每一个元素的选择，选择加入子集或者不加入子集
*/

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		// 每次递归都将当前 path 加入结果
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		// 遍历所有可能的元素
		for i := start; i < len(nums); i++ {
			// 选择当前元素
			path = append(path, nums[i])
			backtrack(i + 1)          // 递归搜索下一个元素
			path = path[:len(path)-1] // 回溯，撤销选择
		}
	}
	backtrack(0)
	return res
}
