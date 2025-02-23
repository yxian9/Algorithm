package leetcode

// @leet start
func validUtf8(data []int) bool {
	var remain int

	for _, b := range data {
		c, mask := 0, 128

		for b&mask != 0 {
			c++
			if c > 4 {
				return false
			}
			mask = mask >> 1

		}

		if c == 0 {
			if remain != 0 {
				return false
			}
			continue
		}

		if c > 1 {
			if remain != 0 {
				return false
			}
			remain = c - 1
			continue
		}

		if remain == 0 {
			return false
		}

		remain--

	}

	return remain == 0
}

// @leet end

