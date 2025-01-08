package leetcode

// @leet start
type pt struct {
	r, c int
}

func (p pt) move(p1 pt) pt {
	return pt{p.r + p1.r, p.c + p1.c}
}

type solution struct {
	input      [][]int
	steps      map[pt]int
	nrow, ncol int
	dirs       []pt
}

func (s *solution) bfs() {
	start := pt{0, 0}
	queue := []pt{start}
	s.steps[start] = 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range s.dirs {
			nextP := cur.move(dir)
			if s.steps[nextP] > 0 {
				continue
			}
			if s.isInside(nextP) && s.get(nextP) == 0 {
				s.steps[nextP] = s.steps[cur] + 1
				queue = append(queue, nextP)
			}
		}
	}
}

func (s *solution) isInside(p pt) bool {
	return p.c >= 0 && p.c < s.ncol && p.r >= 0 && p.r < s.nrow
}

func (s *solution) get(p pt) int {
	return s.input[p.r][p.c]
}

func (s *solution) ans() int {
	res := s.steps[pt{s.nrow - 1, s.ncol - 1}]
	if res == 0 {
		return -1
	}
	return res
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	var dirs []pt
	for _, i := range [3]int{-1, 0, 1} {
		for _, j := range [3]int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			dirs = append(dirs, pt{i, j})
		}
	}
	s := solution{
		input: grid,
		steps: map[pt]int{},
		dirs:  dirs,
		nrow:  len(grid),
		ncol:  len(grid[0]),
	}

	s.bfs()
	return s.ans()
}

// @leet end

