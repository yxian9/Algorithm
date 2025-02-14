package leetcode

type l struct {
	val        int
	prev, next *l
}
type MyLinkedList struct {
	dummy *l
	len   int
}

func Constructor3() MyLinkedList {
	dummy := &l{}
	dummy.next, dummy.prev = dummy, dummy
	return MyLinkedList{dummy, 0}
}

func (this *MyLinkedList) Get(index int) int {
	n := this.getNode(index)
	if n != nil {
		return n.val
	}
	return -1
}

func (this *MyLinkedList) getNode(index int) *l {
	if index >= this.len {
		return nil
	}
	cur := this.dummy.next
	for range index {
		cur = cur.next
	}
	return cur
}

func (this *MyLinkedList) AddBeforeNode(node *l, val int) {
	x := &l{val: val}
	x.prev = node.prev
	x.next = node
	x.prev.next = x
	x.next.prev = x
	this.len++
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddBeforeNode(this.dummy.next, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddBeforeNode(this.dummy, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.len {
		this.AddAtTail(val)
		return
	}
	n := this.getNode(index)
	if n != nil {
		this.AddBeforeNode(n, val)
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	x := this.getNode(index)
	if x != nil {
		x.prev.next = x.next
		x.next.prev = x.prev
		x.next, x.prev = nil, nil
		this.len--
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
