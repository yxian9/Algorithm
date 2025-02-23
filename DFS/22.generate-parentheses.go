package leetcode

import "strings"

// @leet start
func generateParenthesis(n int) []string {
	var (
		ret, path []string
		dfs       func(l, r int)
	)
	dfs = func(l, r int) {
		if l == r && l == n {
			ret = append(ret, strings.Join(path, ""))
			return
		}
		if l < n {
			path = append(path, "(")
			dfs(l+1, r)
			path = path[:len(path)-1]
		}
		if r < l {
			path = append(path, ")")
			dfs(l, r+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ret
}

// @leet end

