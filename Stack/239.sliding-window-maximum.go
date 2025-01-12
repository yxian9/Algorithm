package leetcode

// @leet start
type s239 struct {
	ans    []int
	window []int
}

func (s *s239) push(v int) {
	for len(s.window) > 0 && v > s.window[len(s.window)-1] {
		s.window = s.window[:len(s.window)-1]
	}
	// push only if v <= s.lastItem
	s.window = append(s.window, v)
}

func (s *s239) pop(v int) {
	if len(s.window) > 0 && s.window[0] == v {
		s.window = s.window[1:]
	}
}

func (s *s239) update() {
	s.ans = append(s.ans, s.window[0])
}

func maxSlidingWindow(nums []int, k int) []int {
	s := s239{}
	// fill window for first k elmement
	for i, v := range nums {
		if i == k {
			break
		}
		s.push(v)
	}
	s.update()
	// loop window
	for i := k; i < len(nums); i++ {
		s.push(nums[i])
		s.pop(nums[i-k])
		s.update()
	}
	return s.ans
}

// @leet end

