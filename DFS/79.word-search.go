package leetcode

// @leet start
type Pt struct {
	r, c int
}

// move returns a new point by moving p by the offset d.
func (p Pt) move(d Pt) Pt {
	return Pt{r: p.r + d.r, c: p.c + d.c}
}

func exist(board [][]byte, word string) bool {
	nr := len(board)
	nc := len(board[0])

	dirs := []Pt{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	visited := make(map[Pt]bool)

	// dfs is a recursive helper function that searches for the word.
	var dfs func(p Pt, step int) bool
	dfs = func(p Pt, step int) bool {
		if p.r < 0 || p.r >= nr || p.c < 0 || p.c >= nc {
			return false
		}
		if board[p.r][p.c] != word[step] || visited[p] {
			return false
		}
		if step == len(word)-1 {
			return true
		}

		visited[p] = true

		for _, d := range dirs {
			if dfs(p.move(d), step+1) {
				return true
			}
		}

		visited[p] = false
		return false
	}

	// Try starting DFS from every cell that matches the first letter.
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if board[r][c] == word[0] {
				if dfs(Pt{r, c}, 0) {
					return true
				}
			}
		}
	}
	return false
}

// @leet end

