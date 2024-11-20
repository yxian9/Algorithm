package leetcode

import "slices"

// @leet start
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	var (
		dfs  func(int)
		res  [][]int
		path []int
	)
	n := len(nums)
	dfs = func(idx int) {
		res = append(res, append([]int{}, path...))

		for i := idx; i < n; i++ {
      if i > idx && nums[i] == nums[i-1] {
        continue
      }
			path = append(path, nums[i])
			// dfs(idx + 1)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

// @leet end
