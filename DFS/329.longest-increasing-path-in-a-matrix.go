package leetcode

// @leet start
type pt1 struct {
	r, c int
}

func (p pt1) move(p1 pt1) pt1 {
	return pt1{p.r + p1.r, p.c + p1.c}
}

var dirs1 = [4]pt1{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func longestIncreasingPath(matrix [][]int) int {
	var (
		res  int
		nr   = len(matrix)
		nc   = len(matrix[0])
		memo = map[pt1]int{}
	)

	isInside := func(p pt1) bool {
		return p.r >= 0 && p.r < nr && p.c >= 0 && p.c < nc
	}

	var dfs func(p pt1, prev int) int
	dfs = func(p pt1, prev int) int {
		if !isInside(p) {
			return 0
		}
		curVal := matrix[p.r][p.c]
		if curVal <= prev {
			return 0 // hit the end
		}
		// not simple memo has mean ok. need make sure prev pass the check,
		// then use memo, similar like memo pt + direction.
		// here only point though. Like DP, move on only when avaiable
		if res, ok := memo[p]; ok { // this must check after
			// need make sure current status is valid
			// other wise, will lead to 5>4 then return max length from 4
			return res
		}
		res := 0
		for _, dir := range dirs1 {
			np := p.move(dir)
			res = max(res, dfs(np, curVal))
		}
		memo[p] = res + 1
		return res + 1
	}
	// traaverse
	for r := range nr {
		for c := range nc {
			res = max(res, dfs(pt1{r, c}, -1))
		}
	}
	return res
}

// @leet end
func longestIncreasingPath2(matrix [][]int) int {
	// time exceed. need the bfs
	var (
		res int
		nr  = len(matrix)
		nc  = len(matrix[0])
	)
	isInside := func(p pt1) bool {
		return p.r >= 0 && p.r < nr && p.c >= 0 && p.c < nc
	}
	var dfs func(p pt1, level, prev int)
	dfs = func(p pt1, level, prev int) {
		if !isInside(p) {
			return
		}
		curVal := matrix[p.r][p.c]
		if !(curVal > prev) {
			return
		}
		res = max(res, level)
		for _, dir := range dirs1 {
			nP := p.move(dir)
			dfs(nP, level+1, curVal)
		}
	}
	// traaverse
	for r := range nr {
		for c := range nc {
			dfs(pt1{r, c}, 1, -1)
		}
	}
	return res
}

