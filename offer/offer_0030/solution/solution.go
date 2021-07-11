package solution

type MinStack struct {
	stackData []int
	stackMin  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stackData = append(this.stackData, x)
	if len(this.stackMin) == 0 {
		this.stackMin = append(this.stackMin, x)
	} else {
		if x <= this.Min() {
			this.stackMin = append(this.stackMin, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stackData) == 0 {
		panic("your stack is empty")
	}
	index := len(this.stackData) - 1
	value := this.stackData[index]
	this.stackData = this.stackData[:index]
	if value == this.Min() {
		this.stackMin = this.stackMin[:len(this.stackMin)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stackData) == 0 {
		panic("your stack is empty")
	}
	return this.stackData[len(this.stackData)-1]
}

func (this *MinStack) Min() int {
	if len(this.stackMin) == 0 {
		panic("your stack is empty")
	}
	return this.stackMin[len(this.stackMin)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
