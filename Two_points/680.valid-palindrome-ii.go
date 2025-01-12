package leetcode

// @leet start
func valid(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return valid(s, l+1, r) || valid(s, l, r-1)
		}
		l++
		r--
	}

	return true
}

// @leet end

