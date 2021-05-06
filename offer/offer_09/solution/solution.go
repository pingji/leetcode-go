package solution

type CQueue struct {
	stackPush []int
	stackPop  []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stackPush = append(this.stackPush, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackPop) == 0 {
		if len(this.stackPush) == 0 {
			return -1
		}
		for len(this.stackPush) > 0 {
			index := len(this.stackPush) - 1
			value := this.stackPush[index]
			this.stackPush = this.stackPush[:index]
			this.stackPop = append(this.stackPop, value)
		}
	}
	index := len(this.stackPop) - 1
	value := this.stackPop[index]
	this.stackPop = this.stackPop[:index]
	return value
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
