package solution

type KthLargest struct {
	data     []int
	size     int
	capacity int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		data:     []int{0},
		size:     0,
		capacity: k,
	}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.size < this.capacity {
		this.size++
		this.data = append(this.data, val)
		this.ShiftUp(this.size)
	} else if val > this.data[1] {
		this.data[1] = val
		this.ShiftDown(1)
	}
	return this.data[1]
}

func (this *KthLargest) ShiftUp(i int) {
	for i > 1 && this.data[i/2] > this.data[i] {
		j := i / 2
		this.data[i], this.data[j] = this.data[j], this.data[i]
		i = j
	}
}

func (this *KthLargest) ShiftDown(i int) {
	for 2*i <= this.size {
		j := i * 2
		if j+1 <= this.size && this.data[j+1] < this.data[j] {
			j += 1
		}
		if this.data[i] <= this.data[j] {
			break
		}
		this.data[i], this.data[j] = this.data[j], this.data[i]
		i = j
	}

}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
