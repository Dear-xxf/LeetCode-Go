package GreedyAlgorithm

/*
   题目描述：
	Middle 划分字母区间
	给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
	注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
	返回一个表示每个字符串片段的长度的列表。
*/

/*
   解题思路：
	贪心算法；
	开辟一个数组记录字符串中所有字母最后出现的位置；另外用end记录此次字符串区间最近必须到哪里；
	当遍历数组index等于end的时候，说明一个区间划分结束了；更新start和end，开始一个新的区间。
*/

func partitionLabels(s string) []int {
	res := []int{}
	last := [26]int{}
	for index, value := range s {
		last[value-'a'] = index
	}
	start, end := 0, 0
	for index, value := range s {
		if end < last[value-'a'] {
			end = last[value-'a']
		}
		if index == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
