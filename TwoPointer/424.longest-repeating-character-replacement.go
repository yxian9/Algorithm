package leetcode

// @leet start
func characterReplacement(s string, k int) int {
	var (
		freq            = map[byte]int{}
		maxFreq, res, l int
	)
	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		maxFreq = max(maxFreq, freq[s[r]])

		for r-l+1-maxFreq > k { // NOTE: greed
			freq[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

// @leet end

