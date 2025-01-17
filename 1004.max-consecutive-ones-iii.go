package leetcode

// @leet start
func longestOnes(nums []int, k int) int {
	var l, ret, cnt0 int
	for r, v := range nums {
		cnt0 += 1 - v
		// if 0 exceed k count
		for cnt0 > k {
			cnt0 -= 1 - nums[l]
			l++
		}

		ret = max(ret, r-l+1)
	}
	return ret
}

// @leet end

