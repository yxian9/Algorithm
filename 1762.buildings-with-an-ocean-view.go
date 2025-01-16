package leetcode

import "slices"

// @leet start
func findBuildings(heights []int) []int {
	curMax := 0
	res := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > curMax {
			res = append(res, i)
			curMax = heights[i]
		}
	}
	slices.Sort(res)
	return res
}

// @leet end

