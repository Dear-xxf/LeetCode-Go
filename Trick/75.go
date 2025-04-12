package Trick

/*
   题目描述：
	Middle 颜色分类
	给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
	我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
	必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

/*
   解题思路：
	荷兰国旗问题；
	维护三个区域：
		nums[0..low-1] 是所有的 0；
		nums[low..mid-1] 是所有的 1；
		nums[high+1..n-1] 是所有的 2；
		nums[mid..high] 是未处理区域。
	指针定义：
		low：下一个 0 应该出现的位置；
		mid：当前正在扫描的位置；
		high：下一个 2 应该出现的位置。
	遍历逻辑：
	当 mid <= high 时，检查 nums[mid]：
	如果是 0：交换 nums[low] 和 nums[mid]；
		low++，mid++；
	如果是 1：
		mid++；
	如果是 2：
		交换 nums[mid] 和 nums[high]；
		high--（mid 不动，因为交换过来的值还没处理）
*/

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
