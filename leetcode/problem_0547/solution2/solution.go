package solution

func findCircleNum(isConnected [][]int) (ans int) {
	length := len(isConnected)
	if length == 0 {
		ans = 0
		return
	}
	uf := NewUnionFind(length)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.count
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
