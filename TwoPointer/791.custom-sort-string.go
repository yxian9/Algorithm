package leetcode

import "strings"

// @leet start
func customSortString(order string, s string) string {
	var (
		res  strings.Builder
		freq = [128]int{}
	)
	for _, r := range s {
		freq[r]++
	}
	for _, r := range order {
		for range freq[r] {
			res.WriteRune(r)
		}
		freq[r] = 0
	}
	for i, count := range freq {
		for range count {
			res.WriteRune(rune(i))
		}
	}
	return res.String()
}

// @leet end

