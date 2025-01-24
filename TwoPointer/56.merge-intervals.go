package leetcode

import (
	"sort"
)

// @leet start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := res[len(res)-1]
		if last[1] < interval[0] {
			res = append(res, interval)
		} else {
			last[1] = max(last[1], interval[1])
		}
	}
	return res
}

// @leet end

