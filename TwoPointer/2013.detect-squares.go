package leetcode

// @leet start
type DetectSquares struct {
	pm  map[[2]int]int
	xym map[int]map[int]bool
}

func Constructor() DetectSquares {
	return DetectSquares{
		pm:  map[[2]int]int{},
		xym: map[int]map[int]bool{},
	}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	this.pm[[2]int{x, y}]++
	ym := this.xym[x]
	if ym == nil {
		ym = make(map[int]bool)
		this.xym[x] = ym
	}
	ym[y] = true
}

func (this *DetectSquares) Count(point []int) int {
	ret, x, y := 0, point[0], point[1]
	for y2 := range this.xym[x] {
		if y2 == y {
			continue
		}
		newx := [2]int{x + y2 - y, x + y - y2}
		for _, x2 := range newx {
			ret += this.pm[[2]int{x2, y}] * this.pm[[2]int{x2, y2}] * this.pm[[2]int{x, y2}]
		}
	}
	return ret
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
// @leet end

