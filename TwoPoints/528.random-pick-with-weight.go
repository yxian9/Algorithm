package leetcode

import "math/rand"

// @leet start
type Solution struct {
	total int
	psum  []int
}

func Constructor(w []int) Solution {
	total := 0
	psum := make([]int, len(w))
	for i, v := range w {
		total += v
		psum[i] = total
	}
	return Solution{
		total: total,
		psum:  psum,
	}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.total) + 1
	// find idx in presum
	l, r, idx := 0, len(this.psum)-1, -1
	for l <= r {
		mid := l + (r-l)>>1
		if this.psum[mid] >= target {
			idx = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @leet end

