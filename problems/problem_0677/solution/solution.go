package solution

type TrieNode struct {
	end    bool
	value  int
	childs map[rune]*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		end:    false,
		value:  0,
		childs: make(map[rune]*TrieNode),
	}
}

type MapSum struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		root: newTrieNode(),
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	for _, v := range key {
		if _, ok := cur.childs[v]; !ok {
			cur.childs[v] = newTrieNode()
		}
		cur = cur.childs[v]
	}
	cur.end = true
	cur.value = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for _, v := range prefix {
		if _, ok := cur.childs[v]; !ok {
			return 0
		}
		cur = cur.childs[v]
	}
	return calSum(cur)
}

func calSum(node *TrieNode) int {
	sum := 0
	if node.end {
		sum += node.value
	}
	for _, v := range node.childs {
		sum += calSum(v)
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
