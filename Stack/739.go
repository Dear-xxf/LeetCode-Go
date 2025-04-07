package Stack

/*
   题目描述：
	Middle 每日温度
	给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
	其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
*/

/*
   解题思路：
	熟悉单调栈的话这题很好解决，单调栈主要用于解决下一个更大/更小的问题；
	核心在于使用单调栈记录尚未被“解决”的元素，这里是没有找到下一个更大的日期；
	在遍历的每一个循环中，当我们发现当前日期的温度大于栈顶元素的时候我们就弹出栈顶元素；
	重复这个步骤，直到栈为空或者栈顶元素大于当前日期温度；
	但是go语言这段代码执行时间和内存占比都不算很理想，不确定为什么，可能与读取栈和temperatures中元素的方式有关。
*/

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := []int{}
	for index, value := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < value {
			answer[stack[len(stack)-1]] = index - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	return answer
}
