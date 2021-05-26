package solution

import (
	"container/heap"
)

type Word struct {
	Str   string
	Count int
}

type WordHeap []*Word

func (h WordHeap) Len() int {
	return len(h)
}

func (h WordHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Str < h[j].Str
	}
	return h[i].Count > h[j].Count
}

func (h WordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WordHeap) Push(x interface{}) {
	*h = append(*h, x.(*Word))
}

func (h *WordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

func topKFrequent(words []string, k int) []string {
	var p = make(map[string]int)
	for _, v := range words {
		p[v]++
	}

	var h WordHeap
	for s, c := range p {
		w := &Word{
			Str:   s,
			Count: c,
		}
		heap.Push(&h, w)
	}
	heap.Init(&h)

	res := []string{}
	for i := 0; i < k; i++ {
		v := heap.Pop(&h).(*Word)
		heap.Fix(&h, 0)
		res = append(res, v.Str)
	}

	return res
}
