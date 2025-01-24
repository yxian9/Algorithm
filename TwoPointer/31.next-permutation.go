package leetcode

import "slices"

// @leet start
func nextPermutation(nums []int) {
	n := len(nums)
	p := -1

	// Step 1: Find the pivot
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			p = i
			break
		}
	}

	// If no pivot, reverse the array
	if p == -1 {
		slices.Reverse(nums)
		return
	}

	// Step 2: Find the successor and swap
	for i := n - 1; i > p; i-- {
		if nums[i] > nums[p] {
			nums[i], nums[p] = nums[p], nums[i]
			break
		}
	}

	// Step 3: Reverse the suffix
	for l, r := p+1, n-1; l < r; {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// @leet end

