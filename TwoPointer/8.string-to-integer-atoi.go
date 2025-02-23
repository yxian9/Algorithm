package leetcode

import (
	"math"
	"unicode"
)

// @leet start
func myAtoi(s string) int {
	var (
		num, i int
		sign   = 1
	)
	for i < len(s) && s[i] == ' ' {
		i++
	}
	s = s[i:]
	if len(s) == 0 {
		return 0
	}
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		sign = -1
		s = s[1:]
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			break
		}
		num = num*10 + int(r-'0')
		if sign*num <= math.MinInt32 {
			return math.MinInt32
		}
		if sign*num >= math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return num * sign
}

// @leet end

