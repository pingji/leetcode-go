package solution

type MyQueue struct {
	stackPush []int
	stackPop  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackPush = append(this.stackPush, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	value := this.Peek()
	index := len(this.stackPop) - 1
	this.stackPop = this.stackPop[:index]
	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackPop) == 0 {
		for len(this.stackPush) > 0 {
			index := len(this.stackPush) - 1
			value := this.stackPush[index]
			this.stackPush = this.stackPush[:index]
			this.stackPop = append(this.stackPop, value)
		}
	}
	index := len(this.stackPop) - 1
	value := this.stackPop[index]
	return value
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stackPush) == 0 && len(this.stackPop) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
