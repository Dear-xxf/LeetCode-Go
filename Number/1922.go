package Number

/*
   题目描述：
	Middle 统计好数字的数目
	我们称一个数字字符串是 好数字 当它满足（下标从 0 开始）偶数 下标处的数字为 偶数 且 奇数 下标处的数字为 质数 （2，3，5 或 7）。
	比方说，"2582" 是好数字，因为偶数下标处的数字（2 和 8）是偶数且奇数下标处的数字（5 和 2）为质数。但 "3245" 不是 好数字，因为 3 在偶数下标处但不是偶数。
	给你一个整数 n ，请你返回长度为 n 且为好数字的数字字符串 总数 。由于答案可能会很大，请你将它对 109 + 7 取余后返回 。
	一个 数字字符串 是每一位都由 0 到 9 组成的字符串，且可能包含前导 0 。
*/

/*
   解题思路：
	实现的思路很简单，乘法原理；
	这道题主要考的是代码的实现，要抛弃go语言自带的pow(float64)函数，而去手写快速幂，用平方来加速幂运算。
*/

const mod = 1e9 + 7

func pow(x, n int) int {
	res := 1
	x %= mod
	for n > 0 {
		if n%2 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res
}

func countGoodNumbers(n int64) int {
	evenCount := (n + 1) / 2
	oddCount := n / 2
	return pow(5, int(evenCount)) * pow(4, int(oddCount)) % mod
}
