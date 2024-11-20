package leetcode

import "slices"

// @leet start
func combinationSum2(candidates []int, target int) [][]int {
  slices.Sort(candidates) 
	var (
		res  [][]int
		path []int
		dfs  func(int, int)
		n    int
	)
	n = len(candidates)

	dfs = func(idx, total int) {
		if idx == n || total >= target {
			if total == target {
				res = append(res, append(make([]int, 0, len(path)), path...))
			}
			return
		}
		for i := idx; i < n; i++ {
			if i > idx && candidates[i-1] == candidates[i] {
				continue
			}
			v := candidates[i]
			path = append(path, v)
			dfs(i+1, total+v)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}

// @leet end

