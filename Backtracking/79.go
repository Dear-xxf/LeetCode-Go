package Backtracking

/*
   题目描述：
	Middle 单词搜索
	给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
	单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

/*
   解题思路：
	思路还是比较好想的；两点没有做好；
	第一点是used数组的初始化没有做好；第二点是要关注到单词的开头不一定在[0][0]位置，所以要遍历数组的元素，一一作为开头
	剩下使用回溯的方法不是很困难
*/

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var backtrack func(index, pos int) bool
	backtrack = func(index, pos int) bool {
		row := pos / len(board[0])
		col := pos % len(board[0])
		if word[index] != board[row][col] {
			return false
		}
		if index == len(word)-1 {
			return true
		}

		used[row][col] = true
		for _, value := range direction {
			newRow := row + value[0]
			newCol := col + value[1]
			// 没有超出矩阵范围
			if newRow > len(board)-1 || newRow < 0 || newCol > len(board[0])-1 || newCol < 0 {
				continue
			}
			// 这个位置没有被使用过
			if used[newRow][newCol] {
				continue
			}
			if backtrack(index+1, newRow*len(board[0])+newCol) {
				return true
			}
		}
		used[row][col] = false
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(0, i*len(board[0])+j) {
				return true
			}
		}
	}
	return false
}
