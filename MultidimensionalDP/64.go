package MultidimensionalDP

/*
   题目描述：
	Middle 最小路径和
	给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
	说明：每次只能向下或者向右移动一步。
*/

/*
   解题思路：
	与62思路类似；注意不另外创建dp切片，在原数组上直接操作会省下很多内存。
*/

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row-1][col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
