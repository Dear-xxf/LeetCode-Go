package Matrix

/*
	题目描述：
	Middle 旋转图像
	给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
	你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
*/

/*
	解题思路：
	解题思路与Q54螺旋矩阵类似，将矩阵视为一层套一层的数据结构
	旋转只需要对每一层分别处理即可
	比较麻烦的仍然是边界条件和旋转之前之后的索引
*/

func Rotate(matrix [][]int) {
	n := len(matrix)
	xmin, xmax := 0, n-1
	ymin, ymax := 0, n-1
	for xmin < xmax && ymin < ymax {
		for i := xmin; i < xmax; i++ {
			// 三个临时变量记录右边 下边 左边的对应数据
			temp1 := matrix[i][ymax]
			temp2 := matrix[xmax][ymax-i+ymin]
			temp3 := matrix[xmax-i+xmin][ymin]
			matrix[i][ymax] = matrix[xmin][i] // 用上边的数据覆盖右边的数据
			matrix[xmax][ymax-i+ymin] = temp1 // 填充下边的数据
			matrix[xmax-i+xmin][ymin] = temp2 // 填充右边的数据
			matrix[xmin][i] = temp3           // 填充上边的数据
		}
		xmin++
		xmax--
		ymin++
		ymax--
	}
}
