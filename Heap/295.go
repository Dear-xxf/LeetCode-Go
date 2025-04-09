package Heap

/*
   题目描述：
	Hard 数据流的中位数
	中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
	例如 arr = [2,3,4] 的中位数是 3 。
	例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
	实现 MedianFinder 类:
	MedianFinder() 初始化 MedianFinder 对象。
	void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
	double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。
*/

/*
   解题思路：
	使用两个堆，一个小顶堆，一个大顶堆，各自保存一半的元素;
	大顶堆存放较小的哪一部分，小顶堆存放较大的那一部分；总数为奇数的话放在大顶堆中；
	在添加元素的时候，注意平衡两个堆。
*/

type MedianFinder struct {
	minHeap     []int
	maxHeap     []int
	minHeapSize int
	maxHeapSize int
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap:     make([]int, 0),
		minHeap:     make([]int, 0),
		maxHeapSize: 0,
		minHeapSize: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeapSize == 0 || num <= this.maxHeap[0] {
		addNodeToMaxHeap(&this.maxHeap, num)
		this.maxHeapSize++
	} else {
		addNodeToMinHeap(&this.minHeap, num)
		this.minHeapSize++
	}

	// 平衡两个堆
	if this.maxHeapSize > this.minHeapSize+1 {
		top := extractTopFromMaxHeap(&this.maxHeap)
		this.maxHeapSize--
		addNodeToMinHeap(&this.minHeap, top)
		this.minHeapSize++
	} else if this.minHeapSize > this.maxHeapSize {
		top := extractTopFromMinHeap(&this.minHeap)
		this.minHeapSize--
		addNodeToMaxHeap(&this.maxHeap, top)
		this.maxHeapSize++
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeapSize > this.minHeapSize {
		return float64(this.maxHeap[0])
	}
	return float64(this.maxHeap[0]+this.minHeap[0]) / 2.0
}

func addNodeToMaxHeap(nums *[]int, node int) {
	*nums = append(*nums, node)
	child := len(*nums) - 1
	parent := (child - 1) / 2
	for parent >= 0 {
		if (*nums)[parent] >= (*nums)[child] {
			break
		}
		(*nums)[parent], (*nums)[child] = (*nums)[child], (*nums)[parent]
		child = parent
		parent = (parent - 1) / 2
	}
}

func addNodeToMinHeap(nums *[]int, node int) {
	*nums = append(*nums, node)
	child := len(*nums) - 1
	parent := (child - 1) / 2
	for parent >= 0 {
		if (*nums)[parent] <= (*nums)[child] {
			break
		}
		(*nums)[parent], (*nums)[child] = (*nums)[child], (*nums)[parent]
		child = parent
		parent = (parent - 1) / 2
	}
}

func extractTopFromMaxHeap(nums *[]int) int {
	top := (*nums)[0]
	last := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	if len(*nums) == 0 {
		return top
	}
	(*nums)[0] = last
	heapifyDownMax(nums, 0)
	return top
}

func extractTopFromMinHeap(nums *[]int) int {
	top := (*nums)[0]
	last := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	if len(*nums) == 0 {
		return top
	}
	(*nums)[0] = last
	heapifyDownMin(nums, 0)
	return top
}

func heapifyDownMax(nums *[]int, index int) {
	n := len(*nums)
	for {
		largest := index
		left := index*2 + 1
		right := index*2 + 2
		if left < n && (*nums)[left] > (*nums)[largest] {
			largest = left
		}
		if right < n && (*nums)[right] > (*nums)[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		(*nums)[index], (*nums)[largest] = (*nums)[largest], (*nums)[index]
		index = largest
	}
}

func heapifyDownMin(nums *[]int, index int) {
	n := len(*nums)
	for {
		smallest := index
		left := index*2 + 1
		right := index*2 + 2
		if left < n && (*nums)[left] < (*nums)[smallest] {
			smallest = left
		}
		if right < n && (*nums)[right] < (*nums)[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		(*nums)[index], (*nums)[smallest] = (*nums)[smallest], (*nums)[index]
		index = smallest
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
