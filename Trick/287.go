package Trick

/*
   题目描述：
	Middle 寻找重复数
	给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
	假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
	你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
*/

/*
   解题思路：
	将数字放到对应位置的方法不多解释；
	快慢指针方法可以不修改原数组；
	道理是将这个数组视为一串链表，Next指向为nums[i]；
	那么寻找重复数字的过程就是寻找链表中环的入口，实现方法与Q142 环形链表Ⅱ完全一致。
*/

func findDuplicate(nums []int) int {
	// 快慢指针法
	slow := nums[0]
	fast := nums[0]
	// 第一步：快慢指针找到相遇点
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	// 第二步：找到环的入口，即重复的数字
	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
	// 修改了nums数组
	//temp := 0
	//for index, value := range nums {
	//	for index != value-1 {
	//		if nums[value-1] == value {
	//			return value
	//		}
	//		temp = nums[value-1]
	//		nums[value-1] = value
	//		nums[index] = temp
	//	}
	//}
	//return 0
}
