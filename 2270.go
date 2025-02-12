package leetcode

func waysToSplitArray(nums []int) int {
	num := 0
	leftSum := 0
	rightSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum -= nums[i]
		if leftSum >= rightSum {
			num++
		}
	}
	return num
}
