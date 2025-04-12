package Trick

/*
   题目描述：
	Easy 多数元素
	给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

/*
   解题思路：
	哈希表，一次遍历即可。
*/

func majorityElement(nums []int) int {
	hashTable := map[int]int{}
	for _, value := range nums {
		hashTable[value]++
		if hashTable[value] > len(nums)/2 {
			return value
		}
	}
	return 0
	// 摩尔投票法
	//candidate, count := 0, 0
	//for _, num := range nums {
	//	if count == 0 {
	//		candidate = num
	//		count = 1
	//	} else if num == candidate {
	//		count++
	//	} else {
	//		count--
	//	}
	//}
	//return candidate
}
