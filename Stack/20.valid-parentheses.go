package leetcode

// @leet start
func isValid(s string) bool {
	var (
		stack []rune
		dict  = map[rune]rune{
			')': '(',
			'}': '{',
			']': '[',
		}
	)
	for _, r := range s {
		if res, ok := dict[r]; !ok {
			// find left pair enqueue.
			stack = append(stack, r)
		} else { // try cancel pair
			n := len(stack)
			if n == 0 { // no item
				return false
			}
			if res != stack[n-1] {
				return false
			}
			stack = stack[:n-1] // pop
		}
	}
	return len(stack) == 0
}

// @leet end

