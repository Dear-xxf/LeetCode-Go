package Heap

/*
   题目描述：
	Middle 前K个高频元素
	给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
*/

/*
   解题思路：
	也是比较考验实现堆基本功的一道题；
	思路：
	统计频率
		使用哈希表 map[int]int，遍历数组 nums，统计每个元素出现的次数。
	构建最大堆（优先队列）
		把每个元素和它的频率封装成一个结构体 HeapNode{num, freq}，放入一个最大堆中，频率越高的节点越在上面。
	取出前 K 个元素
		每次从堆顶取出一个元素（最大频率），重复 K 次，将这些元素的 num 加入结果数组。
*/

type HeapNode struct {
	num  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	fre := []HeapNode{}
	hashTable := map[int]int{}
	for _, value := range nums {
		hashTable[value]++
	}
	for index, value := range hashTable {
		addNode(&fre, HeapNode{num: index, freq: value})
	}
	for i := 0; i < k-1; i++ {
		res = append(res, fre[0].num)
		deleteTop(&fre)
	}
	return res
}

func addNode(nums *[]HeapNode, node HeapNode) {
	*nums = append(*nums, node)
	child := len(*nums) - 1
	parent := (child - 1) / 2
	for parent >= 0 {
		if (*nums)[parent].freq >= (*nums)[child].freq {
			break
		}
		(*nums)[parent], (*nums)[child] = (*nums)[child], (*nums)[parent]
		child = parent
		parent = (parent - 1) / 2
	}
}

func deleteTop(nums *[]HeapNode) {
	(*nums)[0], (*nums)[len(*nums)-1] = (*nums)[len(*nums)-1], (*nums)[0]
	*nums = (*nums)[:len(*nums)-1]
	parent := 0
	for {
		left, right := parent*2+1, parent*2+2
		largest := parent
		if left < len(*nums) && (*nums)[left].freq > (*nums)[largest].freq {
			largest = left
		}
		if right < len(*nums) && (*nums)[right].freq > (*nums)[largest].freq {
			largest = right
		}
		if largest == parent {
			break
		}
		(*nums)[parent], (*nums)[largest] = (*nums)[largest], (*nums)[parent]
		parent = largest
	}
}
