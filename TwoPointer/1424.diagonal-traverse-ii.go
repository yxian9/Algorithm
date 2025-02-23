package leetcode

import "slices"

// @leet start
func findDiagonalOrder(nums [][]int) []int {
	m := map[int][]int{}
	var n int
	for i, row := range nums {
		for j := 0; j < len(row); j++ {
			m[i+j] = append(m[i+j], row[j])
			n = max(n, i+j)
		}
	}
	res := make([]int, 0, n)
	for i := range n + 1 {
		cur := m[i]
		slices.Reverse(cur)
		res = append(res, cur...)
	}
	return res
}

// @leet end

