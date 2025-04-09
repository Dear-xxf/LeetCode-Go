package GreedyAlgorithm

import "math"

/*
   题目描述：
	Easy 买卖股票的最佳时机
	给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
	你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
	返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/

/*
   解题思路：
	记录历史最低点；在遍历数组时，判断当前最小值和当天价格之间的利润是多少；
	同时更新最小值和利润。
*/

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	profit := 0
	for _, value := range prices {
		if value < minPrice {
			minPrice = value
		} else {
			if profit < value-minPrice {
				profit = value - minPrice
			}
		}
	}
	return profit
}
