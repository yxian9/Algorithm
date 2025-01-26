package leetcode

// @leet start
type p2 struct {
	r, c int
}

func (p p2) move(p1 p2) p2 {
	return p2{p.r + p1.r, p.c + p1.c}
}

func maxAreaOfIsland(grid [][]int) int {
	var (
		nr, nc = len(grid), len(grid[0])
		dirs   = [4]p2{
			{0, 1},
			{0, -1},
			{-1, 0},
			{1, 0},
		}
		res int
	)
	isInside := func(p p2) bool {
		return p.r >= 0 && p.r < nr && p.c >= 0 && p.c < nc
	}

	var dfs func(p p2) int
	dfs = func(p p2) int {
		if !isInside(p) || grid[p.r][p.c] != 1 {
			return 0
		}
		grid[p.r][p.c] = 0
		res := 0
		for _, dir := range dirs {
			res += dfs(p.move(dir))
		}
		return res + 1
	}

	for r := range nr {
		for c := range nc {
			p := p2{r, c}
			if grid[r][c] == 0 {
				continue
			}
			res = max(res, dfs(p))
		}
	}
	return res
}

// @leet end

