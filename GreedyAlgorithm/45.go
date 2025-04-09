package GreedyAlgorithm

/*
   题目描述：
	Middle 跳跃游戏Ⅱ
	给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
	每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
	0 <= j <= nums[i]
	i + j < n
	返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
*/

/*
   解题思路：
	暴力解法太过耗时；优化思路为使用贪心算法，每一次贪心寻找下一次条约能达到的最远距离；
	我们从左向右遍历 nums 数组，维护两个变量：
		end：当前一步跳跃所能到达的最远位置
		farthest：从当前位置开始，下一步跳跃所能到达的最远位置
	每当我们走到 end 时，必须进行一次跳跃，并更新 end = farthest
*/

func jump(nums []int) int {
	// 暴力，太过耗时
	//hashTable := map[int]int{}
	//hashTable[0] = 0
	//for index, value := range nums {
	//	for i := 1; i <= value; i++ {
	//		if index+i >= len(nums) {
	//			break
	//		}
	//		if hashTable[index+i] > hashTable[index]+1 || hashTable[index+i] == 0 {
	//			hashTable[index+i] = hashTable[index] + 1
	//		}
	//	}
	//}
	//return hashTable[len(nums)-1]

	steps := 0    // 跳跃次数
	end := 0      // 当前能到达的最远边界
	farthest := 0 // 下一步能到达的最远位置
	for i := 0; i < len(nums)-1; i++ {
		// 更新最远位置
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// 到达当前边界，必须跳跃
		if i == end {
			steps++
			end = farthest
		}
	}
	return steps
}
