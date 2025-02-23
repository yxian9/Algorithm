package leetcode

// @leet start
func plusOne(digits []int) []int {
	n := len(digits) - 1
	for ; n >= 0; n-- {
		if digits[n]+1 != 10 {
			digits[n]++
			return digits
		}
		digits[n] = 0
		if n == 0 {
			// ret := []int{1}
			// ret = append(ret, digits...)
			return append([]int{1}, digits...)
			// return ret
		}
	}
	// if n == -1 {
	// 	// ret := []int{1}
	// 	// ret = append(ret, digits...)
	// 	return append([]int{1}, digits...)
	// 	// return ret
	// }
	return digits
}

// @leet end

