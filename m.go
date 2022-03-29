package main

type LFUCache struct {
	keyToVal  map[int]*DLinkNode // key to val 的映射
	keyToFreq map[int]int        // key to freq 的映射
	// freqToKey map[int]*DLinkNode // freq to key 的映射
	freqToKey freqLink
	minFreq   int // 最小频率
	capacity  int // 容量
	size      int // 当前容量
}

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

type freqLink struct {
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

func initDLinkedNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func initFreqLink() freqLink {
	f := freqLink{
		cache: map[int]*DLinkNode{},
		head:  initDLinkedNode(0, 0),
		tail:  initDLinkedNode(0, 0),
	}

	f.head.next = f.tail
	f.tail.prev = f.head

	return f
}

func Constructor(capacity int) LFUCache {
	l := LFUCache{
		capacity:  capacity,
		keyToVal:  map[int]*DLinkNode{},
		freqToKey: initFreqLink(),
		keyToFreq: map[int]int{},
	}

	return l
}

func (this *LFUCache) Get(key int) int {
	if this.size == 0 {
		return -1
	}

	if _, ok := this.keyToVal[key]; !ok {
		return -1
	}

	node := this.keyToVal[key]

	// increase freq
	this.increaseFreq(node)

	return node.value
}

func (this *LFUCache) increaseFreq(node *DLinkNode) {
	// 调整 keyToFreq 当前节点的 freq
	freq := this.keyToFreq[node.key]
	this.keyToFreq[node.key] = freq + 1

	// 从原队列删除
	waitFreq := this.freqToKey.cache[freq]

	this.removeNode(waitFreq)

	// 加入频次更高的队列
	// addNode := this.freqToKey.cache[freq+1]
	this.moveHead(node, freq+1)

}

func (this *LFUCache) moveHead(node *DLinkNode, f int) {
	// freq := this.freqToKey.cache[f]

	node.prev = this.freqToKey.head
	node.next = this.freqToKey.head.next
	this.freqToKey.head.next.prev = node
	this.freqToKey.head.next = node
}

func (this *LFUCache) Put(key int, value int) {
	if _, ok := this.keyToVal[key]; !ok { // 不存在
		// 当前大小是否等于容量
		if this.size >= this.capacity {
			this.removeTail()
		}

		node := &DLinkNode{
			key:   key,
			value: value,
		}

		// set key to val
		this.keyToVal[key] = node

		// set key to freq
		this.keyToFreq[key] = 0

		// increase freq
		this.increaseFreq(node)

		this.size++
		this.minFreq = 1
	} else {
		node := this.keyToVal[key]
		node.value = value
		this.increaseFreq(node)
	}
}

// 移除最小频率最少使用的节点
func (this *LFUCache) removeNode(node *DLinkNode) {
	if node == nil {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LFUCache) removeTail() {
	node := this.freqToKey.cache[this.minFreq]
	this.removeNode(node)

	// 删除 keyToVal 和 keyToFreq 的值
	delete(this.keyToVal, node.key)
	delete(this.keyToFreq, node.key)

	this.size--

	return
}

func main() {
	lrucache := Constructor(2)
	lrucache.Put(1, 1)
}
