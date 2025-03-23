package LinkedList

/*
   题目描述：
	Middle LRU缓存
	请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
	实现 LRUCache 类：
	LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
	int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
	函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

/*
   解题思路：
	双向链表：用于维护元素的访问顺序。每次访问或插入一个元素时，将其移动到链表头部。链表尾部即为最久未使用的元素，当容量超过限制时删除尾部元素。
	哈希表：存储键到链表节点的映射，使得查找操作可以在O(1)时间内完成。
*/

type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LRUNode),
		head:     &LRUNode{},
		tail:     &LRUNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &LRUNode{key: key, value: value}
		this.cache[key] = newNode
		this.addToHead(newNode)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LRUNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
