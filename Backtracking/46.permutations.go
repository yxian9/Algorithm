package leetcode

// @leet start
func permute(nums []int) [][]int {
	n := len(nums)
	res, path, seen := make([][]int, 0), make([]int, n), make([]bool, n)

	var dfs func(int)

	dfs = func(idx int) {
		if idx == n {
			res = append(res, append([]int{}, path...))
			return
		}
		// for i, v := range nums{ // this i will shaow the global i
		for i, v := range nums { //
			if seen[i] {
				continue
			}
			// path[i] = v need to change idx, to i
			path[idx] = v
			seen[i] = true
			dfs(idx + 1)
			seen[i] = false
		}
	}
	dfs(0)
	return res
}

// @leet end
