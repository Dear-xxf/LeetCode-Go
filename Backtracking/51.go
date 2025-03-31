package Backtracking

/*
   题目描述：
	Hard N皇后
	按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
	n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
	给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
	每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

/*
   解题思路：
	经典N皇后问题，实现并不复杂；
	以row为搜索对象，每一次向后搜索下一行摆放情况；
	值得注意的是两个对角线的map在比较时是怎么做的：row - j和row + j；
	并且遇到了一个语法问题，字符串是不能改变的，在Go语言中怎么处理；
*/

func solveNQueens(n int) [][]string {
	results := [][]string{}
	res := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		res = append(res, string(row))
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	var backtrack func(row int, columns, diagonals1, diagonals2 map[int]bool)
	backtrack = func(row int, columns, diagonals1, diagonals2 map[int]bool) {
		if row == n {
			temp := make([]string, len(res))
			copy(temp, res)
			results = append(results, temp)
		}
		// 遍历列
		for j := 0; j < n; j++ {
			if columns[j] {
				continue
			}
			diagonal1 := row - j
			if diagonals1[diagonal1] {
				continue
			}
			diagonal2 := row + j
			if diagonals2[diagonal2] {
				continue
			}

			temp := []byte(res[row])
			temp[j] = byte('Q')
			res[row] = string(temp)

			columns[j] = true
			diagonals1[diagonal1], diagonals2[diagonal2] = true, true
			backtrack(row+1, columns, diagonals1, diagonals2)

			temp = []byte(res[row])
			temp[j] = byte('.')
			res[row] = string(temp)

			delete(columns, j)
			delete(diagonals1, diagonal1)
			delete(diagonals2, diagonal2)
		}
	}
	backtrack(0, columns, diagonals1, diagonals2)
	return results
}
