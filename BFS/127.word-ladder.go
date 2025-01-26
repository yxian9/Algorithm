package leetcode

// @leet start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		queue = []string{beginWord}
		dict  = map[string]bool{}
		step  int
	)
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[endWord] {
		return 0
	}

	for len(queue) > 0 {
		step++
		for range len(queue) {
			w := queue[0]
			queue = queue[1:]
			if w == endWord {
				return step
			}
			// enqueue all the posibile 1 diff word
			for i := 0; i < len(w); i++ {
				for char := 'a'; char <= 'z'; char++ {
					newW := w[:i] + string(char) + w[i+1:]
					if dict[newW] {
						queue = append(queue, newW)
						dict[newW] = false
					}
				}
			}
		}

	}

	return 0
}

// @leet end

