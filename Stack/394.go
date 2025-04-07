package Stack

import (
	"strconv"
	"strings"
)

/*
   题目描述：
	Middle 字符串解码
	给定一个经过编码的字符串，返回它解码后的字符串。
	编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
	你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
	此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
*/

/*
   解题思路：
	个人觉得非常考验基本功的一题，思路并不是很难想，主要是在处理各种字符串符号上比较麻烦
	涉及到多个数字转化为字符串，然后是字符串翻转等小细节；
	具体的实现思路不难，如果遇到数字，将一连串的数字生成一个字符串；
	遇到字母逐个存储；遇到']'处理之间栈内内容直到出现'['；
	最后将整个栈内内容转化为字符串。
*/

func decodeString(s string) string {
	stack := []string{}
	p := 0
	for p < len(s) {
		cur := s[p]
		if cur >= '0' && cur <= '9' {
			// 将p的地址传入，在getDigits函数中修改
			digits := getDigits(s, &p)
			stack = append(stack, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stack = append(stack, string(cur))
			p++
		} else {
			p++
			temp := []string{}
			for stack[len(stack)-1] != "[" {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 翻转string类型切片
			for i := 0; i < len(temp)/2; i++ {
				temp[i], temp[len(temp)-i-1] = temp[len(temp)-i-1], temp[i]
			}
			// 去除'['
			stack = stack[:len(stack)-1]
			digits, _ := strconv.Atoi(stack[len(stack)-1])
			// 去除数字
			stack = stack[:len(stack)-1]
			t := strings.Repeat(getString(temp), digits)
			stack = append(stack, t)
		}
	}
	return getString(stack)
}

// 获取数字，因为数字可能连续出现
func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
