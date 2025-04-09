package GreedyAlgorithm

/*
   题目描述：
	Middle 跳跃游戏
	给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
*/

/*
   解题思路：
	贪心算法；
*/

func canJump(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	farest := 0
	index := 0
	for index <= farest {
		if farest < index+nums[index] {
			farest = index + nums[index]
			if farest >= len(nums)-1 {
				return true
			}
		} else {
			break
		}
		index++
	}
	return false
}
