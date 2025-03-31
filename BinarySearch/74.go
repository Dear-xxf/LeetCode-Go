package BinarySearch

/*
   题目描述：
	Middle 搜索二维矩阵
	给你一个满足下述两条属性的 m x n 整数矩阵：
	每行中的整数从左到右按非严格递增顺序排列。
	每行的第一个整数大于前一行的最后一个整数。
	给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
*/

/*
   解题思路：
	两次搜索，第一次搜索第一列，找到最相近且小于target的元素；第二次在该行上搜索
	这里实现的是z字形搜索，从右上角开始；如果小于target值，行数加1；如果大于target的值，列数减1。
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row != len(matrix) && col != -1 && matrix[row][col] != target {
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return !(row == len(matrix) || col == -1)
}
