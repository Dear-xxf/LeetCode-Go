package NormalArray

/*
	题目描述：
	Middle 合并区间
	以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

/*
	解题思路：
	排序，这道题最核心的问题是排序，只要排序正确，问题非常好解决，遍历即可
*/

func merge(intervals [][]int) [][]int {
	intervals = quickSort(intervals)
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		cur := intervals[i]
		if cur[0] <= last[len(last)-1] && cur[len(cur)-1] > last[len(last)-1] {
			res = res[:len(res)-1]
			res = append(res, []int{last[0], cur[len(cur)-1]})
		}
		if cur[0] > last[len(last)-1] {
			res = append(res, cur)
		}
	}
	return res
}

func quickSort(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	left := [][]int{}
	right := [][]int{}
	middle := [][]int{}
	pivot := intervals[len(intervals)/2][0]
	for _, v := range intervals {
		if v[0] == pivot {
			middle = append(middle, v)
		} else if v[0] > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, middle...), right...)
}
