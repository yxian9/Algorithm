package leetcode

// @leet start
func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var sema int
	for i, r := range s {
		if locked[i] == '0' || r == '(' {
			sema++
		} else if r == ')' {
			sema--
		}
		if sema < 0 {
			return false
		}
	}
	sema = 0
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		if locked[i] == '0' || r == ')' {
			sema++
		} else if r == '(' {
			sema--
		}
		if sema < 0 {
			return false
		}
	}
	return true
}

// @leet end

