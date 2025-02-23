package leetcode

// @leet start
func shortestBridge(grid [][]int) int {
	type pt = [2]int
	var (
		step int
		q    []pt
		n    = len(grid)
		dirs = [4]pt{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		}
		dfs      func(pt)
		isInside = func(p pt) bool {
			return p[0] >= 0 && p[0] < n && p[1] >= 0 && p[1] < n
		}
		get  = func(p pt) int { return grid[p[0]][p[1]] }
		set  = func(p pt, v int) { grid[p[0]][p[1]] = v }
		next = func(p, d pt) pt { return pt{p[0] + d[0], p[1] + d[1]} }
	)
	dfs = func(p pt) {
		if !isInside(p) || get(p) != 1 {
			return
		}
		set(p, 2)
		q = append(q, p)
		for _, d := range dirs {
			dfs(next(p, d))
		}
	}
top:
	for r := range n {
		for c := range n {
			if grid[r][c] == 1 {
				dfs(pt{r, c})
				break top
			}
		}
	}

	for len(q) > 0 {
		for range len(q) {
			cur := q[0]
			q = q[1:]
			for _, d := range dirs {
				np := next(cur, d)
				if isInside(np) {
					if get(np) == 1 {
						return step
					}
					if get(np) == 0 {
						set(np, 2)
						q = append(q, np)
					}
				}
			}
		}
		step++
	}

	return step
}

// @leet end

