package GroupTheory

/*
   题目描述：
	Middle 岛屿数量
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
	岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
	此外，你可以假设该网格的四条边均被水包围。
*/

/*
   解题思路：
	1.DFS;遇到一个‘1’的话，进行DFS，最终岛屿的个数就是进行DFS的次数
	2.并查集；第一时间想到了并查集，太久没用几乎全部忘记了并查集的做法；
	代码如下由gpt生成，借此题熟悉一下go中并查集的做法
*/

// 并查集结构
type UnionFind struct {
	parent []int
	rank   []int
	count  int // 记录连通分量的个数
}

// 初始化并查集
func NewUnionFind(grid [][]byte) *UnionFind {
	rows, cols := len(grid), len(grid[0])
	parent := make([]int, rows*cols)
	rank := make([]int, rows*cols)
	count := 0

	// 初始化并查集，每个陆地是一个独立的集合
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				index := r*cols + c
				parent[index] = index
				count++ // 陆地的初始数量
			}
		}
	}

	return &UnionFind{parent, rank, count}
}

// 查找根节点（带路径压缩）
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// 合并两个集合（按秩合并）
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// 按秩合并
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
		uf.count-- // 每次合并，连通分量减少
	}
}

// 计算岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	uf := NewUnionFind(grid)

	// 方向数组（只向右和向下）
	directions := [][]int{{0, 1}, {1, 0}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				index1 := r*cols + c
				// 遍历右和下的方向，合并相邻陆地
				for _, d := range directions {
					nr, nc := r+d[0], c+d[1]
					if nr < rows && nc < cols && grid[nr][nc] == '1' {
						index2 := nr*cols + nc
						uf.Union(index1, index2)
					}
				}
			}
		}
	}

	return uf.count
}
