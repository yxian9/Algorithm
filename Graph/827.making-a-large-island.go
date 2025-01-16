package leetcode

// @leet start
type s8 struct {
	grid   [][]int
	areaMP map[int]int
	dirs   []p8
	ret    int
}

func (s *s8) set(pt p8, v int) {
	s.grid[pt.r][pt.c] = v
}

func (s *s8) get(pt p8) int {
	return s.grid[pt.r][pt.c]
}

func (s *s8) dfs(p p8, id int) {
	if !s.inside(p) || s.get(p) != 1 {
		return
	}
	s.areaMP[id]++
	s.set(p, id)
	for _, dir := range s.dirs {
		s.dfs(p.move(dir), id)
	}
}

func (s *s8) find(p p8) {
	total := 0
	seen := map[int]bool{}
	for _, dir := range s.dirs {
		next := p.move(dir)
		if !s.inside(next) {
			continue
		}
		id := s.get(next)
		if seen[id] {
			continue
		}
		total += s.areaMP[id]
		seen[id] = true
	}
	s.ret = max(s.ret, total+1)
}

func (s *s8) inside(pt p8) bool {
	nr, nc := len(s.grid), len(s.grid[0])
	return pt.r >= 0 && pt.r < nr && pt.c >= 0 && pt.c < nc
}

func largestIsland(grid [][]int) int {
	sol := s8{
		areaMP: map[int]int{},
		grid:   grid,
		dirs:   dirs,
	}
	id := 2
	for r := range len(grid) {
		for c := range len(grid[0]) {
			curP := p8{r, c}
			if sol.get(curP) == 1 {
				// sol.curArea = 0
				sol.dfs(curP, id)
				// sol.areaMP[id] = sol.curArea
				id++
			}
		}
	}
	found0 := false
	// find max area
	for r := range len(grid) {
		for c := range len(grid[0]) {
			curP := p8{r, c}
			if sol.get(curP) == 0 {
				found0 = true
				sol.find(curP)
			}
		}
	}

	if !found0 {
		return len(grid) * len(grid[0])
	}
	return sol.ret
}

var dirs = []p8{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

type p8 struct {
	r, c int
}

func (p p8) move(p1 p8) p8 {
	return p8{p.r + p1.r, p.c + p1.c}
}

// @leet end

