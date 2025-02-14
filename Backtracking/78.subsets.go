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

func subsets3(nums []int) [][]int {
	res, path, seen := [][]int{}, []int{}, make([]bool, len(nums))
	n := len(nums)
	var dfs func(int)
	dfs = func(idx int) {
		res = append(res, append([]int{}, path...))
		for i := idx; i < n; i++ {
			if seen[i] {
				continue
			}
			path = append(path, nums[i])
			seen[i] = true // use seen, not i+1 to remove order duplicate
			dfs(i)
			path = path[:len(path)-1]
			seen[i] = false
		}
	}
	dfs(0)
	return res
}

// wireld approach
func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		res = append(res, path)
		for i := idx; i < len(nums); i++ {
			temp := []int{}
			temp = append(temp, path...)
			temp = append(temp, nums[i])
			dfs(i+1, temp)
		}
	}
	dfs(0, []int{})
	return res
}
