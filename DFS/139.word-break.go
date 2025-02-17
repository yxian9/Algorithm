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
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	m := map[string]bool{}
	choice := map[int]struct{}{}
	for _, v := range wordDict {
		m[v] = true
		choice[len(v)] = struct{}{}
	}

	for i := range len(s) + 1 {
		for step := range choice {
			if i >= step && dp[i-step] {
				if m[s[i-step:i]] {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)]
}


