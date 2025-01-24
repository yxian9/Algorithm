package leetcode

// @leet start
func isToeplitzMatrix(matrix [][]int) bool {
	nrow, ncol := len(matrix), len(matrix[0])
	check := func(r, c int) bool {
		num := matrix[r][c]
		for r < nrow && c < ncol {
			if matrix[r][c] != num {
				return false
			}
			r++
			c++
		}
		return true
	}
	for r := range nrow {
		if !check(r, 0) {
			return false
		}
	}
	for c := range ncol {
		if !check(0, c) {
			return false
		}
	}
	return true
}

// @leet end

