package leetcode

// @leet start
type MovingAverage struct {
	size, count, sum int
	arr              []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		arr:  make([]int, 0, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.count < this.size {
		this.sum += val
		this.arr = append(this.arr, val)
	} else {
		idx := this.count % this.size
		this.sum = this.sum - this.arr[idx] + val
		this.arr[idx] = val
	}
	this.count++
	return float64(this.sum) / float64(len(this.arr))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
// @leet end

