package NormalArray

import "math"

/*
   题目描述：
	Easy 统计位数为偶数的数字
	给你一个整数数组 nums，请你返回其中包含 偶数 个数位的数字的个数。
*/

/*
   解题思路：
	问题关键在于如何判断一个数字的位数个数；
	两种方法
	一是转化为字符串判断字符串长度；
	二是使用log函数后向下取整，得到 k−1≤log(x)<k，因此我们可以用 k=⌊log(x)+1⌋ 得到 x 包含的数字个数 k。
*/

func findNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		if int(math.Log10(float64(num))+1)%2 == 0 {
			res++
		}
	}
	return res
}
