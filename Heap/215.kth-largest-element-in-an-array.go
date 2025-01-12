package leetcode

import "slices"

// @leet start
func findKthLargest(nums []int, k int) int {
	slices.Sort(nums)
	var res int
	for i := len(nums) - 1; i >= 0; i-- {
		res = nums[i]
		k--
		if k == 0 {
			return res
		}
	}
	return res
}

// @leet end

