package leetcode

import (
	"slices"
	"strconv"
)

// @leet start

func findTargetSumWays(nums []int, target int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	total := 0
	for i := range len(nums) {
		total += nums[i]
	}
	if abs(target) > total {
		return 0
	}
	if remain := (total + target) % 2; remain == 1 {
		return 0
	}
	total = (total + target) / 2
	var dfs func(int, int) int
	slices.Sort(nums)
	dfs = func(idx, cur int) int {
		if cur == total {
			// compared to fire and go, this cur == total will prune the
			// branch, even with sorting, [0,0] target 0, expect 4, which will only report 1
			return 1
		}
		res := 0
		for i := idx; i < len(nums); i++ {
			res += dfs(i+1, cur+nums[i])
		}
		return res
	}
	return dfs(0, 0)
}

// @leet end

func findTargetSumWays4(nums []int, target int) int {
	// dfs search
	total := 0
	for i := range len(nums) {
		total += nums[i]
	}
	if abs(target) > total {
		return 0
	}
	fori
	if remain := (total + target) % 2; remain == 1 {
		return 0
	}
	total = (total + target) / 2
	res := 0
	var dfs func(int, int)

	dfs = func(idx, cur int) {
		if cur == total {
			res += 1
		}
		for i := idx; i < len(nums); i++ {
			dfs(i+1, cur+nums[i])
		}
	}
	dfs(0, 0)

	return res
}

func findTargetSumWays3(nums []int, target int) int {
	memo := make(map[string]int)
	var dfs func(int, int) int
	dfs = func(level, total int) int {
		key := strconv.Itoa(level) + ":" + strconv.Itoa(total)
		if v, ok := memo[key]; ok {
			return v
		}
		if level == len(nums) {
			if total == target {
				return 1
			} else {
				return 0
			}
		}

		left := dfs(level+1, total+nums[level])  // select for add
		right := dfs(level+1, total-nums[level]) // select for minus
		memo[key] = left + right
		return left + right
	}

	return dfs(0, 0)
}

func findTargetSumWays2(nums []int, target int) int {
	total := 0
	for i := range len(nums) {
		total += nums[i]
	}
	if abs(target) > total {
		return 0
	}
	if remain := (total + target) % 2; remain == 1 {
		return 0
	}
	total = (total + target) / 2
	dp := make([]int, total+1)
	dp[0] = 1
	// permutation, loop j first, can not repick, j--
	// !! not permutation,  -++ and +-+ different, but not mean permutation
	// combination, loop item first, can not repick j--
	for _, v := range nums {
		for j := len(dp) - 1; j >= 0; j-- {
			if j-v < 0 {
				continue
			}
			if dp[j-v] > 0 {
				dp[j] += dp[j-v]
			}
		}
	}
	// for each item, value == weight. pay attention to the difference
	return dp[len(dp)-1]
}
