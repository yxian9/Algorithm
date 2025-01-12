package leetcode

// @leet start
type node struct {
	Val       int
	Pre, Next *node
}
type MyLinkedList struct {
	length int
	dummy  *node
}

func Constructor() MyLinkedList {
	dummy := &node{}
	dummy.Pre = dummy
	dummy.Next = dummy
	return MyLinkedList{
		dummy:  dummy,
		length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	x := this.getNode(index)
	if x == nil {
		return -1
	}
	return x.Val
}

func (this *MyLinkedList) getNode(index int) *node {
	if index > this.length-1 {
		return nil
	}
	cur := this.dummy.Next
	for range index {
		cur = cur.Next
	}
	return cur
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.inSertBefore(this.dummy.Next, val)
}

func (this *MyLinkedList) inSertBefore(x *node, Val int) {
	nn := &node{Val: Val}
	nn.Pre = x.Pre
	nn.Next = x
	nn.Pre.Next = nn
	nn.Next.Pre = nn
	this.length++
}

func (this *MyLinkedList) deleteNode(x *node) {
	x.Pre.Next = x.Next
	x.Next.Pre = x.Pre
	this.length--
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.inSertBefore(this.dummy, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.length {
		this.AddAtTail(val)
		return
	}
	curNode := this.getNode(index)
	if curNode != nil {
		this.inSertBefore(curNode, val)
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	curNode := this.getNode(index)
	if curNode != nil {
		this.deleteNode(curNode)
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @leet end

