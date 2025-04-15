package NormalArray

/*
   题目描述：
	Hard 统计数组中好三元组数目
	给你两个下标从 0 开始且长度为 n 的整数数组 nums1 和 nums2 ，两者都是 [0, 1, ..., n - 1] 的 排列 。
	好三元组 指的是 3 个 互不相同 的值，且它们在数组 nums1 和 nums2 中出现顺序保持一致。换句话说，如果我们将 pos1v 记为值 v 在 nums1 中出现的位置，pos2v 为值 v 在 nums2 中的位置，
	那么一个好三元组定义为 0 <= x, y, z <= n - 1 ，且 pos1x < pos1y < pos1z 和 pos2x < pos2y < pos2z 都成立的 (x, y, z) 。
	请你返回好三元组的 总数目 。
*/

/*
   解题思路：
	树状数组；
	第一步：映射顺序
		由于两个数组都是0 到 n-1 的排列：
		用一个 pos 数组记录 nums2 中每个元素的位置；
		然后将 nums1 中的值用 pos[val] 映射成在 nums2 中出现的位置。
		这样就转换成了一个问题：在一个数组中，找三元组 (i, j, k)，满足 i < j < k，且 A[i] < A[j] < A[k]。
		这个问题本质上就是：统计数组中递增的三元组。
	第二步：统计递增三元组
		枚举中间位置 j，然后：
		左边统计有多少个 i 满足 A[i] < A[j]；
		右边统计有多少个 k 满足 A[k] > A[j]；
		然后将两边的结果乘起来，就是以 j 为中心的好三元组数量。
		这个过程中，为了快速统计“某个数左边比它小的个数”或“右边比它大的个数”，就可以用树状数组或线段树来做前缀和统计。
*/

type BIT struct {
	tree []int
	size int
}

func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+2),
		size: n + 2,
	}
}

func (b *BIT) update(i int, delta int) {
	i++
	for i < b.size {
		b.tree[i] += delta
		i += i & -i
	}
}

func (b *BIT) query(i int) int {
	i++
	res := 0
	for i > 0 {
		res += b.tree[i]
		i -= i & -i
	}
	return res
}

func goodTriplets(nums1, nums2 []int) int64 {
	n := len(nums1)
	pos := make([]int, n)
	for i, v := range nums2 {
		pos[v] = i
	}

	// order[i] 表示 nums1[i] 在 nums2 中的位置
	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = pos[nums1[i]]
	}

	bit := NewBIT(n)
	leftSmaller := make([]int, n)

	// 从左往右统计左边比当前 order[i] 小的数量
	for i := 0; i < n; i++ {
		leftSmaller[i] = bit.query(order[i] - 1)
		bit.update(order[i], 1)
	}

	// 清空 BIT，用来统计右边比当前大的数量
	bit = NewBIT(n)
	rightLarger := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		rightLarger[i] = bit.query(n-1) - bit.query(order[i])
		bit.update(order[i], 1)
	}

	var result int64
	for i := 0; i < n; i++ {
		result += int64(leftSmaller[i]) * int64(rightLarger[i])
	}
	return result
}
