package DynamicProgramming

/*
   题目描述：
	Easy 杨辉三角
	给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
*/

/*
   解题思路：
	也是经典的DP题目；不多写思路了。
*/

func generate(numRows int) [][]int {
	res := [][]int{{1}}
	for index := 1; index < numRows; index++ {
		lastRow := res[len(res)-1]
		row := make([]int, index+1)
		row[0], row[index] = 1, 1
		for col := 1; col < index; col++ {
			row[col] = lastRow[col-1] + lastRow[col]
		}
		temp := []int{}
		copy(temp, row)
		res = append(res, temp)
	}
	return res
}
