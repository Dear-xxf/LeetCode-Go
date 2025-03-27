package GroupTheory

/*
   题目描述：
	Middle 腐烂的橘子
	在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
	值 0 代表空单元格；
	值 1 代表新鲜橘子；
	值 2 代表腐烂的橘子。
	每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
	返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
*/

/*
   解题思路：
	多源广度优先搜索；
	知道使用BFS之后，这道题目并不是很难；在初始化的时候将所有的腐烂橘子都加入队列中，同时用一个depth哈希表记录每一个橘子腐烂的时间
	也就是BFS搜索的轮次；depth中最大的值就是要返回的分钟数
	当最后还有新鲜橘子的时候，返回-1
*/

func orangesRotting(grid [][]int) int {
	minute := 0
	row, col := len(grid), len(grid[0])
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	num_of_fresh := 0
	list_of_bad := []int{}
	depth := make(map[int]int)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				num_of_fresh++
			}
			if grid[i][j] == 2 {
				list_of_bad = append(list_of_bad, i*col+j)
			}
		}
	}
	for len(list_of_bad) > 0 {
		orange := list_of_bad[0]
		list_of_bad = list_of_bad[1:]
		r, c := orange/col, orange%col
		for k := 0; k < 4; k++ {
			nr, nc := r+dr[k], c+dc[k]
			if nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				num_of_fresh--
				ncode := nr*col + nc
				list_of_bad = append(list_of_bad, ncode)
				depth[ncode] = depth[orange] + 1
				minute = depth[ncode]
			}
		}
	}
	if num_of_fresh == 0 {
		return minute
	} else {
		return -1
	}
}
