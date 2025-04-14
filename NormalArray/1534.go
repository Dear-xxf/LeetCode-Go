package NormalArray

/*
   题目描述：
	Easy 统计好三元组
	给你一个整数数组 arr ，以及 a、b 、c 三个整数。请你统计其中好三元组的数量。
	如果三元组 (arr[i], arr[j], arr[k]) 满足下列全部条件，则认为它是一个 好三元组 。
	0 <= i < j < k < arr.length
	|arr[i] - arr[j]| <= a
	|arr[j] - arr[k]| <= b
	|arr[i] - arr[k]| <= c
	其中 |x| 表示 x 的绝对值。
	返回 好三元组的数量 。
*/

/*
   解题思路：
	暴力+提前剪枝
*/

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	res := 0
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
