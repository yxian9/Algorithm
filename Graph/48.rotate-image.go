package leetcode

// @leet start
func rotate(matrix [][]int) {
	nr, nc := len(matrix), len(matrix[0])
	// transpose
	for i := 0; i < nr; i++ {
		for j := i; j < nc; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	rev := func(row []int) {
		for i, j := 0, len(row)-1; i < j; i++ {
			row[i], row[j] = row[j], row[i]
			j--
		}
	}
	// reverse
	for i := 0; i < nr; i++ {
		// slices.Reverse(matrix[i])
		rev(matrix[i])
	}
}

// @leet end

