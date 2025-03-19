package SlidingWindow

/*
	题目描述：
	Middle 无重复字符的最长子串
	给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串的长度。
*/

/*
	解题思路：
	滑动窗口，当右指针遇到的元素在当前子串中重复的时候
	左指针向右滑动，而右指针不变，因为这种情况下如果存在比当前还长的子串，右指针一定在当前右指针的右边
	使用哈希表记录当前子串中的所有元素
*/

func lengthOfLongestSubstring(s string) int {
	ans := 0
	right := 0
	hashTable := map[byte]int{}
	for left := 0; left < len(s); left++ {
		if left != 0 {
			delete(hashTable, s[left-1])
		}
		for right < len(s) && hashTable[s[right]] == 0 {
			hashTable[s[right]] = 1
			right++
		}
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}
