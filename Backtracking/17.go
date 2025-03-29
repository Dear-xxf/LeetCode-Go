package Backtracking

/*
   题目描述：
	Middle 电话号码的字母组合
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
	给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

/*
   解题思路：
	回溯；与子集类似，对每一位数字进行几种情况的枚举即可
*/

func letterCombinations(digits string) []string {
	res := []string{}
	var phoneMap map[string]string = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var backTrack func(path string)
	backTrack = func(path string) {
		if len(path) == len(digits) {
			if len(path) > 0 {
				res = append(res, path)
			}
			return
		}
		digit := string(digits[len(path)])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			backTrack(path + string(letters[i]))
		}
	}
	backTrack("")
	return res
}
