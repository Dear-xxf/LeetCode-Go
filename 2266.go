package leetcode

/*
	题目描述：
	为了 打出 一个字母，Alice 需要 按 对应字母 i 次，i 是该字母在这个按键上所处的位置。
	比方说，为了按出字母 's' ，Alice 需要按 '7' 四次。类似的， Alice 需要按 '5' 两次得到字母  'k' 。
	注意，数字 '0' 和 '1' 不映射到任何字母，所以 Alice 不 使用它们。
	但是，由于传输的错误，Bob 没有收到 Alice 打字的字母信息，反而收到了 按键的字符串信息 。
	比方说，Alice 发出的信息为 "bob" ，Bob 将收到字符串 "2266622" 。
	给你一个字符串 pressedKeys ，表示 Bob 收到的字符串，请你返回 Alice 总共可能发出多少种文字信息 。
	由于答案可能很大，将它对 109 + 7 取余 后返回。
*/

/*
	解题思路：
	动态规划
*/

func countTexts(pressedKeys string) int {
	const mod = 1_000_000_007
	n := len(pressedKeys)
	dp := make([]int, n+1)
	dp[0] = 1 // 空字符串的初始状态

	for i := 1; i <= n; i++ {
		// 当前按键
		dp[i] = dp[i-1] // 每个字符至少可以单独作为一个字母

		// 检查连续的按键
		for j := 2; j <= 4; j++ {
			if i-j < 0 || pressedKeys[i-1] != pressedKeys[i-j] {
				break
			}
			// 如果是 '7' 或 '9'，最多 4 次；否则最多 3 次
			if j <= 3 || (pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9') {
				dp[i] = (dp[i] + dp[i-j]) % mod
			}
		}
	}

	return dp[n]
}
