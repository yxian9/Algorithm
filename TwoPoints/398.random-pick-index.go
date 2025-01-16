package leetcode

import (
	"math/rand"
	"sort"
)

// @leet start
type i1 struct {
	idx, v int
}
type Solution struct {
	arr []i1
}

func Constructor(nums []int) Solution {
	arr := make([]i1, len(nums))
	for i, v := range nums {
		arr[i] = i1{i, v}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].v < arr[j].v
	})
	return Solution{
		arr: arr,
	}
}

func (this *Solution) left(target int) int {
	l, r, idx := 0, len(this.arr)-1, -1
	for l <= r {
		mid := l + (r-l)>>1
		if this.arr[mid].v >= target {
			idx = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

func (this *Solution) right(target int) int {
	l, r, idx := 0, len(this.arr)-1, -1
	for l <= r {
		mid := l + (r-l)>>1
		if this.arr[mid].v <= target {
			idx = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return idx
}

func (this *Solution) Pick(target int) int {
	lower, upper := this.left(target), this.right(target)
	if upper != lower {
		return this.arr[lower+rand.Intn(upper-lower+1)].idx
	}
	return this.arr[lower].idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @leet end

