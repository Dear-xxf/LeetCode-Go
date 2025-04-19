package NormalArray

/*
   题目描述：
	Middle 统计坏数对的数目
	给你一个下标从 0 开始的整数数组 nums 。如果 i < j 且 j - i != nums[j] - nums[i] ，那么我们称 (i, j) 是一个 坏数对 。
	请你返回 nums 中 坏数对 的总数目。
*/

/*
   解题思路：
	观察到j - i != nums[j] - nums[i]可以转化为i - nums[i] != j - nums[j];
	将问题转化为坏数对的总数目为数对的总数目减去好数对的数目。
*/

func countBadPairs(nums []int) int64 {
	n := len(nums)
	totalPairs := int64(n) * int64(n-1) / 2
	counter := make(map[int]int)

	for i, num := range nums {
		key := i - num
		counter[key]++
	}

	var goodPairs int64 = 0
	for _, cnt := range counter {
		goodPairs += int64(cnt) * int64(cnt-1) / 2
	}

	return totalPairs - goodPairs
}
