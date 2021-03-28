package solution

type Trie struct {
	end    bool
	childs [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		if cur.childs[v-'a'] == nil {
			cur.childs[v-'a'] = new(Trie)
		}
		cur = cur.childs[v-'a']
	}
	cur.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if cur.childs[v-'a'] == nil {
			return false
		}
		cur = cur.childs[v-'a']
	}
	return cur.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if cur.childs[v-'a'] == nil {
			return false
		}
		cur = cur.childs[v-'a']
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
