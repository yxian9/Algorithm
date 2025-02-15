package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	nr, nc := len(matrix), len(matrix[0])
	for i := 1; i < nr; i++ {
		for j := 1; j < nc; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}


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

