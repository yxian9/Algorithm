package leetcode

import "slices"

// @leet start
type s1249 struct {
	chars []rune
	idx   []int
}

func (s *s1249) top() rune {
	return s.chars[len(s.chars)-1]
}

func (s *s1249) pop() {
	s.chars = s.chars[:len(s.chars)-1]
	s.idx = s.idx[:len(s.idx)-1]
}

func (s *s1249) push(r rune, i int) {
	s.chars = append(s.chars, r)
	s.idx = append(s.idx, i)
}

func minRemoveToMakeValid(s string) string {
	solution := s1249{}
	for i, char := range s {
		if char == '(' {
			solution.push(char, i)
		}
		if char == ')' {
			if len(solution.chars) > 0 && solution.top() == '(' {
				solution.pop()
			} else {
				solution.push(char, i)
			}
		}
	}
	res := make([]rune, 0, len(s))
	for idx, char := range s {
		if slices.Contains(solution.idx, idx) {
			continue
		}
		res = append(res, char)
	}
	return string(res)
}

// @leet end

