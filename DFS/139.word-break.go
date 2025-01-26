package leetcode

// @leet start
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	memo := map[int]bool{}
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == n {
			return true
		}
		if idx > n {
			return false
		}
		if res, ok := memo[idx]; ok {
			return res
		}
		res := false
		for _, w := range wordDict {
			if idx+len(w) <= n && s[idx:idx+len(w)] == w {
				if dfs(idx + len(w)) {
					res = true
					break
				}
			}
		}
		memo[idx] = res
		return res
	}
	return dfs(0)
}

// @leet end

