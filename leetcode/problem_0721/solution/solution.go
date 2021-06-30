package solution

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	uf := NewUnionFind(len(accounts))
	mailToID := make(map[string]int)

	for i, account := range accounts {
		for _, mail := range account[1:] {
			id, ok := mailToID[mail]
			if !ok {
				mailToID[mail] = i
			} else {
				uf.Union(id, i)
			}
		}
	}

	idMap := make(map[int][]string)
	for mail, id := range mailToID {
		rootID := uf.Find(id)
		_, ok := idMap[rootID]
		if !ok {
			idMap[rootID] = []string{mail}
		} else {
			idMap[rootID] = append(idMap[rootID], mail)
		}
	}

	res := [][]string{}
	for id, mails := range idMap {
		name := accounts[id][0]
		sort.StringSlice(mails).Sort()
		value := append([]string{name}, mails...)
		res = append(res, value)
	}
	return res
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		count:  n,
		parent: make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(i int) int {
	if i != uf.parent[i] {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) {
	iRoot := uf.Find(i)
	jRoot := uf.Find(j)
	if iRoot != jRoot {
		uf.parent[iRoot] = jRoot
		uf.count--
	}
}
