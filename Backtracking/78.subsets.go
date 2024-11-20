package leetcode

// @leet start
func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	n := len(nums)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			res = append(res, append([]int{}, path...))
			return
		}

		dfs(idx + 1)

		path = append(path, nums[idx])
		dfs(idx + 1)
		path = path[:len(path)-1]
	}
	dfs(0)

	return res
}

// @leet end

func subsets1(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	n := len(nums)
	var dfs func(int)
	dfs = func(idx int) {
		res = append(res, append([]int{}, path...))
		for start := idx; start < n; start++ {
			path = append(path, nums[start])
			dfs(start + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)

	return res
}

