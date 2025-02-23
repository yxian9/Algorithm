package leetcode

import (
	"strconv"
	"strings"
)

// @leet start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	var (
		i   int
		ret strings.Builder
	)
	for j := range len(s) {
		if s[j] != s[i] {
			ret.WriteString(strconv.Itoa(j - i))
			ret.WriteByte(s[i])
			i = j
		}
	}
	ret.WriteByte(byte(len(s)-i) + '0')
	ret.WriteByte(s[len(s)-1])

	return ret.String()
}

// @leet end

