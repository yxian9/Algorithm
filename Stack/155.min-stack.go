package leetcode

import "math"

// @leet start
type item1 struct {
	min, val int
	prev     *item1
}
type MinStack struct {
	head *item1 // NOTE: can not use embed here.
	// need head a a refer point
}

func Constructor() MinStack {
	return MinStack{
		head: &item1{
			min: math.MaxInt32,
		},
	}
}

func (this *MinStack) Push(val int) {
	this.head = &item1{
		prev: this.head,
		val:  val,
		min:  min(this.head.min, val),
	}
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.prev
}

func (this *MinStack) Top() int {
	if this.head == nil {
		panic("emty")
	}
	return this.head.val
}

func (this *MinStack) GetMin() int {
	if this.head == nil {
		panic("emty")
	}
	return this.head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @leet end

