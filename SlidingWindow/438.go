package SlidingWindow

/*
	题目描述：
	Middle
	给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
*/

/*
	解题思路：
	固定一个长度为子串的滑动窗口，每一次向右滑动来判断是否符合异位词的要求
	我的实现思路很顺畅，关键提升点在于每一次滑动之后，不能再做一次循环，来判断是否为异位词
	应当注意到每一次滑动只有滑出的元素和新进入的元素对滑动窗口匹配的结果产生影响
	再考虑到只包含小写字母，所以可以将所有小写字母在滑动过程中都进行记录，最后只需要判断匹配个数的小写字母数量是否为26
*/

func findAnagrams(s string, p string) []int {
	ans := []int{}
	if len(s) < len(p) {
		return ans
	}
	left := 0
	right := left + len(p) - 1
	hashTable := map[byte]int{}
	pHashTable := map[byte]int{}
	for i := left; i < len(p); i++ {
		hashTable[s[i]]++
		pHashTable[p[i]]++
	}
	matchCount := 0
	for ch := byte('a'); ch <= 'z'; ch++ {
		if hashTable[ch] == pHashTable[ch] {
			matchCount++
		}
	}
	for left = 0; left < len(s); left++ {
		if left != 0 {
			right++
			if right >= len(s) {
				break
			}
			// 原先的左指针滑出
			hashTable[s[left-1]]--
			if hashTable[s[left-1]] == pHashTable[s[left-1]] {
				matchCount++
			} else if hashTable[s[left-1]] == pHashTable[s[left-1]]-1 {
				matchCount--
			}
			// 右指针右滑，增加新的元素进入窗口
			hashTable[s[right]]++
			if hashTable[s[right]] == pHashTable[s[right]] {
				matchCount++
			} else if hashTable[s[right]] == pHashTable[s[right]]+1 {
				matchCount--
			}
		}
		if matchCount == 26 {
			ans = append(ans, left)
		}
	}
	return ans
}
