package leetcode

import "slices"

// @leet start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return strs[0]
	}
	slices.Sort(strs)
	l, r := strs[0], strs[len(strs)-1]
	for i := range len(l) {
		if l[i] != r[i] {
			return l[0:i]
		}
	}
	return l
}

// @leet end
