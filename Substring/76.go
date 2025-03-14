package Substring

/*
	题目描述：
	hard
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
*/

/*
	解题思路：
	滑动窗口
*/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// 统计 t 中字符出现的次数
	tCount := make(map[byte]int)
	for i := range t {
		tCount[t[i]]++
	}

	// 滑动窗口中的字符统计
	window := make(map[byte]int)

	// 左右指针
	left, right := 0, 0
	// 记录最小覆盖子串的起始位置和长度
	start, minLen := 0, len(s)+1
	// 记录窗口内匹配了 t 的多少个字符
	matchCount := 0
	for right < len(s) {
		// 扩展右边界
		c := s[right]
		if _, exists := tCount[c]; exists {
			window[c]++
			if window[c] == tCount[c] {
				matchCount++
			}
		}
		right++ // 移动右指针

		// 收缩窗口
		for matchCount == len(tCount) {
			// 更新最小窗口
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			// 左指针收缩
			d := s[left]
			if _, exists := tCount[d]; exists {
				if window[d] == tCount[d] {
					matchCount--
				}
				window[d]--
			}
			left++ // 移动左指针
		}
	}
	// 如果没找到符合条件的子串，返回 ""
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
