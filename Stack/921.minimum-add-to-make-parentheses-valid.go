package leetcode

// @leet start
func minAddToMakeValid(s string) int {
	stack := make([]rune, 0, len(s))
	for _, r := range s {
		switch r {
		case '(':
			stack = append(stack, r)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				break
			}
			stack = append(stack, r)
		}
	}
	return len(stack)
}

// @leet end
