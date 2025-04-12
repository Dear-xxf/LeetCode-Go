package Trick

/*
   题目描述：
	Easy 只出现一次的数字
	给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/

/*
   解题思路：
	位运算；非常巧妙。
	同一个数字按位相加两次之后为0；
	借助此题需要开始熟悉位运算的算法题。
*/

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
