package leetcode

import "math/rand"

// @leet start
type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	} else {
		this.m[val] = len(this.arr)
		this.arr = append(this.arr, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.m[val]; ok {
		n := len(this.arr)
		lst := this.arr[n-1]
		this.arr[idx] = lst
		this.m[lst] = idx

		delete(this.m, val)
		this.arr = this.arr[:n-1]
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.arr))
	return this.arr[n]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @leet end

