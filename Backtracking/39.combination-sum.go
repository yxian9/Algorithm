package leetcode

// @leet start
func combinationSum(candidates []int, target int) [][]int {
	n, res, path := len(candidates), [][]int{}, []int{}
	var dfs func(int, int)
	dfs = func(idx, total int) {
		if total == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if idx == n || total > target {
			return
		}
		// interesting. although we donot constrain level,(can repick), still need idx to constran not chose
		dfs(idx+1, total)
		cur := candidates[idx]
		path = append(path, cur)
		dfs(idx, total+cur)
		path = path[:len(path)-1]
	}
	dfs(0, 0)
	return res
}

// @leet end
func combinationSum2(candidates []int, target int) [][]int {
	n, res, path := len(candidates), [][]int{}, []int{}
	var dfs func(int, int)
	dfs = func(idx, total int) {
		if total == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if total > target {
			return
		}
		for i := idx; i < n; i++ {
			cur := candidates[i]
			// total += cur // not work, total need for next for loop
			path = append(path, cur)
			dfs(i, total+cur)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}
