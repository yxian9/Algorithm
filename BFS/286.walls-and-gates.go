package leetcode

import "math"

// @leet start
type pp struct {
	r, c int
}

func (p pp) move(p1 pp) pp {
	return pp{p.r + p1.r, p.c + p1.c}
}

const empty = math.MaxInt32

func wallsAndGates(rooms [][]int) {
	var (
		nr, nc = len(rooms), len(rooms[0])
		queue  []pp
		dirs   = [4]pp{
			{0, 1},
			{0, -1},
			{-1, 0},
			{1, 0},
		}
		step     int
		isInside = func(p pp) bool {
			return p.r >= 0 && p.r < nr && p.c >= 0 && p.c < nc
		}

		bfs = func() {
			for len(queue) > 0 {
				step++
				for range len(queue) {
					lh := queue[0]
					queue = queue[1:]
					// rooms[lh.r][lh.c] = step // this not work
					// when enqueue. mean seen
					// when do post process, first item in queue may step on later item at queu3
					for _, dir := range dirs {
						np := lh.move(dir)
						if isInside(np) && rooms[np.r][np.c] == empty {
							queue = append(queue, np)
							rooms[np.r][np.c] = step
						}
					}
				}
			}
		}
	)

	for r := range nr {
		for c := range nc {
			if rooms[r][c] == 0 {
				queue = append(queue, pp{r, c})
			}
		}
	}
	bfs()
}

// @leet end

