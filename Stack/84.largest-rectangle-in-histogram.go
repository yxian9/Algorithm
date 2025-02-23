package leetcode

// @leet start
func largestRectangleArea(heights []int) int {
	var (
		res   int
		st    []int
		n     = len(heights)
		left  = make([]int, n)
		right = make([]int, n)
	)
	// for i := range n {
	// 	left[i] = -1
	// }
	// for i := range n {
	// 	right[i] = n
	// }
	for i := range n {
		v := heights[i]
		for len(st) > 0 && heights[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		v := heights[i]
		for len(st) > 0 && heights[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	for i := range n {
		res = max(res, heights[i]*(right[i]-left[i]-1))
	}

	return res
}

// @leet end

