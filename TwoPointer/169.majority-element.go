package leetcode

// @leet start
func majorityElement(nums []int) int {
	var (
		count = 1
		ret   = nums[0]
	)
	for _, v := range nums[1:] {
		if v == ret {
			count++
		} else {
			count--
		}
		if count < 0 {
			count, ret = 0, v
		}
	}
	return ret
}

// @leet end

