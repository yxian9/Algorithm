package leetcode

import "strings"

// /.../a/../b/c/../d/./
// @leet start
type solution71 struct {
	ans []string
}

func (s *solution71) parse(path string) {
	var temp string
	for _, char := range path {
		if char == '/' {
			if len(temp) > 0 {
				s.ans = append(s.ans, temp)
				temp = ""
			}
			continue
		}
		temp = temp + string(char)
	}
	if len(temp) > 0 {
		s.ans = append(s.ans, temp)
	}
}

func (s *solution71) push(temp string) {
	switch temp {
	case ".":
	case "..":
		if len(s.ans) > 0 {
			s.ans = s.ans[0 : len(s.ans)-1]
		}
	default:
		s.ans = append(s.ans, temp)
	}
}

func simplifyPath(path string) string {
	solution := solution71{}
	solution.parse(path)
	return "/" + strings.Join(solution.ans, "/")
}

// @leet end

