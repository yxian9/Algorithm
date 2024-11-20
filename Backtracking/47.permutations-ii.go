package leetcode

import "slices"

// @leet start
func permuteUnique(nums []int) [][]int {
	var (
		dfs  func(int)
		seen []bool
		path []int
		res  [][]int
		n    int
	)
	n = len(nums)
	slices.Sort(nums)
	seen = make([]bool, n)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for j, v := range nums {
			if seen[j] {
				continue
			}
			// if j > 0 && nums[j] == nums[j-1] && seen[j] { j will not work. need check with previous one
			// need to look back
			if j > 0 && nums[j] == nums[j-1] && seen[j-1] {
				continue
			}
			seen[j] = true
			path = append(path, v)
			dfs(i + 1)
			path = path[:len(path)-1]
			seen[j] = false

		}
	}
	dfs(0)
	return res
}

// @leet end

