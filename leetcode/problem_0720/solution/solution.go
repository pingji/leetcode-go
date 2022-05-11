package solution

type Trie struct {
	end    bool
	childs [26]*Trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		if cur.childs[v-'a'] == nil {
			cur.childs[v-'a'] = &Trie{}
		}
		cur = cur.childs[v-'a']
	}
	cur.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if cur.childs[v-'a'] == nil || !cur.childs[v-'a'].end {
			return false
		}
		cur = cur.childs[v-'a']
	}
	return cur.end
}

func longestWord(words []string) (longest string) {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(longest) ||
			(len(word) == len(longest) && word < longest)) {
			longest = word
		}
	}
	return longest
}
