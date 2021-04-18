package solution

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func NewDLinkedNode(key int, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     NewDLinkedNode(0, 0), // dummy head
		tail:     NewDLinkedNode(0, 0), // dummy tail
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	} else {
		node := this.cache[key]
		this.moveToHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := NewDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
