package Matrix

/*
	题目描述：
	Middle 搜索二维矩阵
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。
*/

/*
	解题思路：
	1.可以利用每一行都是顺序排列的性质，对每一行进行二分查找
	2.从右上角开始查找
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	nx := len(matrix)
	ny := len(matrix[0])
	row, col := 0, ny-1
	for row < nx && col > -1 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}
