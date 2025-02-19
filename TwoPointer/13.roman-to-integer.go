package leetcode

// @leet start
func romanToInt(s string) int {
	var (
		m   = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
		ret int
	)
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if i+1 < len(s) && m[cur] < m[s[i+1]] {
			ret -= m[cur]
		} else {
			ret += m[cur]
		}

	}
	return ret
}

// @leet end

