package leetcode

type n struct {
	key, val   int
	next, prev *n
}
type LRUCache struct {
	capacity int
	m        map[int]*n
	dummy    *n
}

func Constructor146(capacity int) LRUCache {
	dummy := &n{}
	dummy.next, dummy.prev = dummy, dummy
	return LRUCache{capacity, map[int]*n{}, dummy}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.detach(v)
		this.leftAppend(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.val = value
		this.detach(v)
		this.leftAppend(v)
		return
	}
	newN := &n{key: key, val: value}
	this.m[key] = newN
	this.leftAppend(newN)
	if len(this.m) > this.capacity {
		lastItem := this.dummy.prev
		this.detach(lastItem)
		delete(this.m, lastItem.key)
	}
}

func (this *LRUCache) detach(x *n) {
	x.prev.next = x.next
	x.next.prev = x.prev
	x.next, x.prev = nil, nil
}

func (this *LRUCache) leftAppend(x *n) {
	x.prev = this.dummy
	x.next = this.dummy.next
	x.prev.next = x
	x.next.prev = x
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

