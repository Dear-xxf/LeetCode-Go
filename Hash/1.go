package Hash

/*
	题目描述：
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
	你可以按任意顺序返回答案。
*/

/*
	解题思路：
	1.暴力解法
	2.哈希表
*/

func twoSum(nums []int, target int) []int {
	// 暴力解法
	//slice := []int{}
	//for i := 0; i < len(nums); i++ {
	//	firstNum := nums[i]
	//	for j := i; j < len(nums); j++ {
	//		secNum := nums[j]
	//		if firstNum+secNum == target {
	//			slice = append(slice, i, j)
	//			return slice
	//		} else {
	//			continue
	//		}
	//	}
	//}
	//return slice

	// 哈希表
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
