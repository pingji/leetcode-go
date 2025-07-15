package solution

import (
	"container/heap"
)

// Word 表示单词及其出现次数
type Word struct {
	Str   string // 单词字符串
	Count int    // 出现次数
}

// WordHeap 是一个最大堆，用于存储 Word 指针
// 按照出现次数降序排序，如果次数相同则按字典序升序排序
type WordHeap []*Word

// Len 返回堆中元素的数量
// 使用值接收者：只读操作，不修改切片
func (h WordHeap) Len() int {
	return len(h)
}

// Less 定义堆的比较逻辑
// 使用值接收者：只是比较元素，不修改切片
// 优先级：出现次数高的在前，次数相同时字典序小的在前
func (h WordHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Str < h[j].Str // 次数相同时，字典序小的优先级高
	}
	return h[i].Count > h[j].Count // 次数高的优先级高
}

// Swap 交换堆中两个元素的位置
// 使用值接收者：虽然是值接收者，但切片是引用类型，
// 可以通过索引修改底层数组的元素，不改变切片本身的结构
func (h WordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 向堆中添加一个元素
// 使用指针接收者：因为需要修改切片的长度（append操作）
// 如果使用值接收者，append后的结果不会影响原切片
func (h *WordHeap) Push(x any) {
	*h = append(*h, x.(*Word))
}

// Pop 从堆中移除并返回优先级最高的元素
// 使用指针接收者：因为需要修改切片的长度
// 通过 *h = old[0 : n-1] 来缩短切片
func (h *WordHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

// topKFrequent 返回给定单词列表中出现频率最高的前 k 个单词
// 参数：
//
//	words: 单词列表
//	k: 需要返回的前 k 个单词数量
//
// 返回：按出现频率降序排列的前 k 个单词，频率相同时按字典序升序排列
func topKFrequent(words []string, k int) []string {
	// 统计每个单词的出现次数
	var p = make(map[string]int)
	for _, v := range words {
		p[v]++
	}

	// 创建最大堆，存储所有单词及其出现次数
	var h WordHeap
	for s, c := range p {
		w := &Word{
			Str:   s,
			Count: c,
		}
		heap.Push(&h, w)
	}
	heap.Init(&h) // 初始化堆

	// 从堆中取出前 k 个单词
	res := []string{}
	for i := 0; i < k; i++ {
		v := heap.Pop(&h).(*Word)
		res = append(res, v.Str)
	}

	return res
}
