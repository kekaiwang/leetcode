/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	inQueue, outQueue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.outQueue = append(this.outQueue, x)

	for len(this.inQueue) > 0 {
		this.outQueue = append(this.outQueue, this.inQueue[0])
		this.inQueue = this.inQueue[1:]
	}

	this.inQueue, this.outQueue = this.outQueue, this.inQueue
}

func (this *MyStack) Pop() int {
	x := this.inQueue[0]
	this.inQueue = this.inQueue[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.inQueue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.inQueue) == 0 && len(this.outQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

