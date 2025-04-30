package GreedyAlgorithm

import (
	"container/list"
	"sort"
)

/*
   题目描述：
	Hard 你可以安排的最多任务数目
	给你 n 个任务和 m 个工人。每个任务需要一定的力量值才能完成，需要的力量值保存在下标从 0 开始的整数数组 tasks 中，第 i 个任务需要 tasks[i] 的力量才能完成。每个工人的力量值保存在下标从 0 开始的整数数组 workers 中，第 j 个工人的力量值为 workers[j] 。每个工人只能完成 一个 任务，且力量值需要 大于等于 该任务的力量要求值（即 workers[j] >= tasks[i] ）。
	除此以外，你还有 pills 个神奇药丸，可以给 一个工人的力量值 增加 strength 。你可以决定给哪些工人使用药丸，但每个工人 最多 只能使用 一片 药丸。
	给你下标从 0 开始的整数数组tasks 和 workers 以及两个整数 pills 和 strength ，请你返回 最多 有多少个任务可以被完成。
*/

/*
   解题思路：
	（好难的贪心算法题，给gpt老师都干冒烟了也没写对）
	最后也是在gpt的帮助下才基本完全掌握了贪心算法+二分查找的解法；
	排序预处理
		任务从小到大排序
		工人从小到大排序
	二分答案
		猜最大能完成的任务数k（范围0到min(任务数,工人数)）
		对每个k值，验证是否能分配成功
	贪心验证（关键）
		选最大的k个任务，按从大到小处理
		选最强的k个工人，按从小到大处理
		对每个任务：
		从最难的任务向简单任务遍历，对于这个任务，优先选择能力较强的工人；
		如果能力最强的工人无法胜任，转而选择嗑药之后能力最弱的工人（且能胜任任务）来完成任务；
	我的初始版本代码在每一次判断canAssign的时候初始状态将所有的mid个工人加入双端队列，
	但是这样做会在处理较为困难的任务时导致误删掉一些能力不够的工人导致后续的其他任务人手不够，所以一个值得注意的点是需要针对每一个任务都要向队列中放置工人。
*/

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	n, m := len(tasks), len(workers)

	canAssign := func(mid int) bool {
		p := pills       // 当前可用药丸数量
		ws := list.New() // 双端队列，保存当前可以尝试分配的工人
		ptr := m - 1     // 从最强的工人开始向前遍历
		// 从最难的任务开始向前尝试分配
		for i := mid - 1; i >= 0; i-- {
			// 将所有有可能完成当前任务的工人放入双端队列
			for ptr >= m-mid && workers[ptr]+strength >= tasks[i] {
				ws.PushFront(workers[ptr])
				ptr--
			}
			// 如果当前没有可用的工人了，任务分配失败
			if ws.Len() == 0 {
				return false
			}
			// 如果工人中最强的可以完成任务，直接用
			if ws.Back().Value.(int) >= tasks[i] {
				ws.Remove(ws.Back())
			} else {
				// 否则尝试给最弱的工人使用一颗药丸
				if p == 0 {
					return false
				}
				p--
				ws.Remove(ws.Front())
			}
		}
		return true // 所有mid个任务都能完成
	}

	left, right, ans := 1, min(m, n), 0
	for left <= right {
		mid := (left + right) / 2
		if canAssign(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
