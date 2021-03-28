package solution2

type Trie struct {
	end    bool
	childs map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		end:    false,
		childs: make(map[rune]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		if _, ok := cur.childs[v]; !ok {
			trie := Constructor()
			cur.childs[v] = &trie
		}
		cur = cur.childs[v]
	}
	cur.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if _, ok := cur.childs[v]; !ok {
			return false
		}
		cur = cur.childs[v]
	}
	return cur.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if _, ok := cur.childs[v]; !ok {
			return false
		}
		cur = cur.childs[v]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
