package solution

type LFUCache struct {
	capacity int
	minFreq  int
	keyMap   map[int]*DLinkedNode
	freqMap  map[int]*DLinkedNodeList
}
type DLinkedNode struct {
	key   int
	value int
	freq  int
	prev  *DLinkedNode
	next  *DLinkedNode
}

type DLinkedNodeList struct {
	size int
	head *DLinkedNode
	tail *DLinkedNode
}

func newDLinkedNode(key int, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
		freq:  0,
		prev:  nil,
		next:  nil,
	}
}

func newDLinkedNodeList() *DLinkedNodeList {
	l := &DLinkedNodeList{
		size: 0,
		head: newDLinkedNode(0, 0), // dummy head
		tail: newDLinkedNode(0, 0), // dummy tail
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *DLinkedNodeList) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	l.size++
}

func (l *DLinkedNodeList) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *DLinkedNodeList) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		minFreq:  0,
		keyMap:   make(map[int]*DLinkedNode),
		freqMap:  make(map[int]*DLinkedNodeList),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if _, ok := this.keyMap[key]; !ok {
		return -1
	} else {
		node := this.keyMap[key]
		this.updateFreqNode(node)
		return node.value
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if _, ok := this.keyMap[key]; !ok {
		if len(this.keyMap) >= this.capacity {
			this.removeTail()
		}
		node := newDLinkedNode(key, value)
		this.keyMap[key] = node
		this.addFreqNode(node)

	} else {
		node := this.keyMap[key]
		node.value = value
		this.updateFreqNode(node)
	}
}

func (this *LFUCache) updateFreqNode(node *DLinkedNode) {
	this.removeFreqNode(node)
	this.addFreqNode(node)

}
func (this *LFUCache) removeFreqNode(node *DLinkedNode) {
	nodes := this.freqMap[node.freq]
	nodes.removeNode(node)
	if nodes.size == 0 {
		delete(this.freqMap, node.freq)
		if node.freq == this.minFreq {
			this.minFreq = 0
		}
	} else {
		this.freqMap[node.freq] = nodes
	}
}

func (this *LFUCache) addFreqNode(node *DLinkedNode) {
	node.freq++
	var nodes *DLinkedNodeList
	if _, ok := this.freqMap[node.freq]; ok {
		nodes = this.freqMap[node.freq]
	} else {
		nodes = newDLinkedNodeList()
	}
	nodes.addToHead(node)
	this.freqMap[node.freq] = nodes
	if node.freq < this.minFreq || this.minFreq == 0 {
		this.minFreq = node.freq
	}
}

func (this *LFUCache) removeTail() {
	nodes := this.freqMap[this.minFreq]
	node := nodes.removeTail()
	delete(this.keyMap, node.key)

	if nodes.size > 0 {
		this.freqMap[this.minFreq] = nodes
	} else {
		delete(this.freqMap, this.minFreq)
	}
}
