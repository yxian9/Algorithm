package leetcode

// @leet start
func strStr(haystack string, needle string) int {
	res := -1
	if len(needle) > len(haystack) {
		return res
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return res
}

// @leet end

