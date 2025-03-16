package NormalArray

/*
	题目描述：
	Middle 除自身以外数组的乘积
	给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
	题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
	请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/

/*
	解题思路：
	难点在于不使用除法，虽然我不理解为什么不给使用除法
	拿到题目没有太多的思路，只能想到构建一个新的数组，两层循环处理
	可以构建两个数组，分别为L和R，L[i]记录i位置左侧的乘积和，R[i]记录i位置右侧的乘积和
*/

func productExceptSelf(nums []int) []int {
	ans := []int{}
	L := []int{1}
	R := []int{1}
	for i := 1; i < len(nums); i++ {
		L = append(L, nums[i-1]*L[i-1])
	}
	for i := len(nums) - 2; i >= 0; i-- {
		// R = append([]int{nums[i+1] * R[0]}, R...)
		// 上面这样写会存在性能问题，在O(n)循环中进行了O(n^2)的操作
		// 改为下面这种写法
		R = append(R, 0)
		copy(R[1:], R)
		R[0] = nums[i+1] * R[1]
	}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, L[i]*R[i])
	}
	return ans
}
