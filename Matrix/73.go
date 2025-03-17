package Matrix

/*
	题目描述：
	Middle 矩阵置零
	给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/

/*
	解题思路：
	想到标记i和j位置并不难，用矩阵原有位置进行标记需要有点技巧
	并且利用矩阵原有位置进行标记的时候需要考虑到[0][0]位置上的特殊情况
	这里让[0][0]位置上的值只负责记录第0行的情况
	第0列另外开辟一个空间col0来记录
*/

func setZeroes(matrix [][]int) {
	col0 := 1 // 记录第一列是否需要置零
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0 = 0 // 记录第一列需要清零
		}
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := len(matrix) - 1; i > -1; i-- {
		for j := len(matrix[0]) - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		// 第一列取出单独处理，因为[0][0]位置上为0会影响到第一列的其他元素的判断
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
}
