package DynamicProgramming

/*
   题目描述：
	Easy 爬楼梯
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

/*
   解题思路：
	动态规划入门题；当前台阶的方法数量等于前一个台阶和前二个台阶两个台阶方法数量的总和；
	优化的点在于可以只使用两个int类型循环记录数组中的内容。
*/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp0 := 1
	dp1 := 2
	for i := 2; i < n; i++ {
		temp := dp1 + dp0
		dp0 = dp1
		dp1 = temp
	}
	return dp1
}
