package Hash

/*
	题目描述：
	给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
	请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

/*
	解题思路：
	哈希
	用哈希表将数组中的数字存储起来，方便快速查找是否存在。
	遍历每个数字，找出只可能作为连续序列起点的数字（num-1 不在哈希表中）。
	从这些起点开始，逐个计算连续序列的长度，并更新最大长度。
	整个过程确保每个数字只被处理一次，时间复杂度为O(n)
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashTable := map[int]bool{}
	for _, num := range nums { // 创建hashTable
		hashTable[num] = true
	}
	maxLength := 0
	for num := range hashTable { // 注意这里只需要读hashTale中的num就可以了
		if !hashTable[num-1] { // 判断num-1是否在hashTable中，防止重复处理
			var cur = num
			curMaxLength := 1
			for hashTable[cur+1] {
				curMaxLength++
				cur++
			}
			if maxLength < curMaxLength {
				maxLength = curMaxLength
			}
		}
	}
	return maxLength
}
