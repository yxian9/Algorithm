package leetcode

// @leet start
func productExceptSelf(nums []int) []int {
	var (
		pre, post = 1, 1
		ret       = make([]int, len(nums))
	)
	for i := range ret {
		ret[i] = 1
	}
	for i := range ret {
		ret[i] *= pre
		pre *= nums[i]
	}
	for i := len(ret) - 1; i >= 0; i-- {
		ret[i] *= post
		post *= nums[i]
	}
	return ret
}

// @leet end

