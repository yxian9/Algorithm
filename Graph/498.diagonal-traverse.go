package leetcode

// @leet start

func findDiagonalOrder(mat [][]int) []int {
	nrow, ncol := len(mat), len(mat[0])
	nums := make([][]int, nrow+ncol-1)
	for r := range nrow {
		for c := range ncol {
			nums[r+c] = append(nums[r+c], mat[r][c])
		}
	}
	res := make([]int, 0, nrow*ncol)
	for i, num := range nums {
		if i%2 == 0 {
		}
		res = append(res, num...)
	}
	return res
}

// @leet end

