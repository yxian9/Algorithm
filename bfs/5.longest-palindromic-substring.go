package leetcode

// @leet start
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i <= len(s); i++ {
		res = maxPalid(res, s, i, i)
		res = maxPalid(res, s, i, i+1)
	}
	return res
}

func maxPalid(cur, s string, l, r int) string {
	res := ""
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res = s[l : r+1]
		l--
		r++
	}
	if len(res) > len(cur) {
		return res
	}
	return cur
}

// @leet end
