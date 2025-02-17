package leetcode

// @leet start
func valid(freq [128]int) bool {
	for _, v := range freq {
		if v == 2 {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	var (
		l, ret int
		freq   [128]int
	)
	for r, v := range s {
		freq[v]++
		for !valid(freq) {
			freq[s[l]]--
			l++
		}
		ret = max(ret, r-l+1)
	}
	return ret
}

// @leet end

