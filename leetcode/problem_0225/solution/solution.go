package solution

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	for len(this.queue1) > 0 {
		value := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, value)
	}
	this.queue1 = append(this.queue1, x)
	for len(this.queue2) > 0 {
		value := this.queue2[0]
		this.queue2 = this.queue2[1:]
		this.queue1 = append(this.queue1, value)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	value := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	value := this.queue1[0]
	return value
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
