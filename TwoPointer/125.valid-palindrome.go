package leetcode

import (
	"unicode"
)

// @leet start
func isPalindrome(s string) bool {
	var (
		l, r = 0, len(s) - 1
		skip = func(r rune) bool {
			return !unicode.IsNumber(r) && !unicode.IsLetter(r)
		}
	)
	for l < r {
		if skip(rune(s[l])) {
			l++
			continue
		}
		if skip(rune(s[r])) {
			r--
			continue
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

// @leet end

