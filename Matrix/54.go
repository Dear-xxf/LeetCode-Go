package Matrix

/*
	题目描述：
	Middle 螺旋矩阵
	给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

/*
	解题思路：
	不算很难的Middle题，将四个方向的情况考虑清楚就没有太大问题
	有一点麻烦的在于需要考虑不能重复，即每一行或者每一列开始位置上的元素是算在上一步的最后还是新步骤的开始
	以及终止条件需要盘算清晰
*/

func spiralOrder(matrix [][]int) []int {
	// 初始化先加入[0][0]位置元素
	res := []int{}
	// direction 0 1 2 3分别代表向右 向下 向左 向上
	direction := 0
	xmin, ymin := 0, 0
	xmax, ymax := len(matrix)-1, len(matrix[0])-1
	for xmin <= xmax && ymin <= ymax {
		switch direction {
		case 0: // 向右 →
			for ycur := ymin; ycur <= ymax; ycur++ {
				res = append(res, matrix[xmin][ycur])
			}
			xmin++
		case 1: // 向下 ↓
			for xcur := xmin; xcur <= xmax; xcur++ {
				res = append(res, matrix[xcur][ymax])
			}
			ymax--
		case 2: // 向左 ←
			for ycur := ymax; ycur >= ymin; ycur-- {
				res = append(res, matrix[xmax][ycur])
			}
			xmax--
		case 3: // 向上 ↑
			for xcur := xmax; xcur >= xmin; xcur-- {
				res = append(res, matrix[xcur][ymin])
			}
			ymin++
		}
		direction = (direction + 1) % 4 // 切换方向
	}
	return res
}
