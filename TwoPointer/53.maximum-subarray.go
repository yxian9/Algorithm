package leetcode

// @leet start
func maxSubArray(nums []int) int {
	cum, res := nums[0], nums[0]
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if cum > 0 {
			cum += v
		} else {
			cum = v
		}
		res = max(res, cum)

	}
	return res
}

// @leet end

