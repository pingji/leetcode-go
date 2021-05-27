package solution

// import "fmt"

type Word struct {
	Str   string
	Count int
}

func (w *Word) LargeThan(w2 *Word) bool {
	if (*w).Count == (*w2).Count {
		return (*w).Str < (*w2).Str
	}
	return (*w).Count > (*w2).Count
}

type WrodMinHeap struct {
	data     []*Word
	size     int
	capacity int
}

func NewWrodMinHeap(k int) WrodMinHeap {
	heap := WrodMinHeap{
		data:     []*Word{nil},
		size:     0,
		capacity: k,
	}
	return heap
}

func (h *WrodMinHeap) Add(w *Word) {
	if h.size < h.capacity {
		h.size++
		h.data = append(h.data, w)
		h.ShiftUp(h.size)
	} else if w.LargeThan(h.data[1]) {
		h.data[1] = w
		h.ShiftDown(1)
	}
}

func (h *WrodMinHeap) Pop() *Word {
	if h.size == 0 {
		return nil
	}
	v := h.data[1]
	if h.size > 1 {
		h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
		h.data = h.data[:h.size]
		h.size--
		h.ShiftDown(1)
	} else {
		h.data = h.data[:h.size]
		h.size--
	}
	return v
}

func (h *WrodMinHeap) ShiftUp(i int) {
	for i > 1 && h.data[i/2].LargeThan(h.data[i]) {
		j := i / 2
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
}

func (h *WrodMinHeap) ShiftDown(i int) {
	for 2*i <= h.size {
		j := i * 2
		if j+1 <= h.size && h.data[j].LargeThan(h.data[j+1]) {
			j += 1
		}
		if h.data[j].LargeThan(h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}

}

func topKFrequent(words []string, k int) []string {
	var p = make(map[string]int)
	for _, v := range words {
		p[v]++
	}

	h := NewWrodMinHeap(k)
	for s, c := range p {
		w := &Word{
			Str:   s,
			Count: c,
		}
		h.Add(w)
	}

	res := []string{}
	for i := 0; i < k; i++ {
		v := h.Pop()
		res = append([]string{v.Str}, res...)
	}
	return res
}

// func (h *WrodMinHeap) print() {
// 	fmt.Printf("heap size: %v, heap capacity: %v\n", h.size, h.capacity)
// 	for i := 1; i < h.size+1; i++ {
// 		fmt.Printf("%v, %v; ", i, *(h.data[i]))
// 	}
// 	fmt.Println("")
// }
