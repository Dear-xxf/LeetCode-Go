package Hash

import "sort"

/*
	题目描述：
	给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
	字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
*/

/*
	解题思路：
	哈希
	先把每一个字符串s转化为切片，对切片进行排序
*/

func groupAnagrams(strs []string) [][]string {
	hashTable := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		hashTable[sortedStr] = append(hashTable[sortedStr], str)
	}
	ans := make([][]string, 0, len(hashTable))
	for _, v := range hashTable {
		ans = append(ans, v)
	}
	return ans
}
