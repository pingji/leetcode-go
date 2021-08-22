package solution

import "container/heap"

type Item struct {
	rowIdx int
	colIdx int
	value  int
}

type ItemHeap []*Item

func (h ItemHeap) Len() int {
	return len(h)
}

func (h ItemHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

func mergeKArrays(nums [][]int) []int {
	rowCnt := len(nums)
	if rowCnt == 0 {
		return []int{}
	}

	res := make([]int, 0)

	var h ItemHeap

	for i := 0; i < rowCnt; i++ {
		item := &Item{
			rowIdx: i,
			colIdx: 0,
			value:  nums[i][0],
		}
		heap.Push(&h, item)
	}

	for h.Len() > 0 {
		v := heap.Pop(&h).(*Item)
		res = append(res, v.value)
		if v.colIdx+1 < len(nums[v.rowIdx]) {
			item := &Item{
				rowIdx: v.rowIdx,
				colIdx: v.colIdx + 1,
				value:  nums[v.rowIdx][v.colIdx+1],
			}
			heap.Push(&h, item)
		}
	}
	return res
}
