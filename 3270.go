package leetcode

func GenerateKey(num1 int, num2 int, num3 int) int {
	num := 0
	multiplier := 1
	for i := 0; i < 4; i++ {
		digit1 := (num1 / multiplier) % 10
		digit2 := (num2 / multiplier) % 10
		digit3 := (num3 / multiplier) % 10
		// 找到当前位的最小值
		curMin := digit1
		if digit2 < curMin {
			curMin = digit2
		}
		if digit3 < curMin {
			curMin = digit3
		}
		// 构造结果数字
		num += curMin * multiplier
		multiplier *= 10 // 更新进位
	}
	return num
}
