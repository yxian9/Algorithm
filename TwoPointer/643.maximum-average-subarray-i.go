package leetcode

// @leet start
func findMaxAverage(nums []int, k int) float64 {
	var total, maxV int
	for i := range k {
		total += nums[i]
	}
	maxV = total
	for i := k; i < len(nums); i++ {
		total += nums[i]
		total -= nums[i-k]
		if total > maxV {
			maxV = total
		}
	}
	return float64(maxV) / float64(k)
}

// @leet end

