package leetcode

// @leet start
func subarraySum(nums []int, k int) int {
	var (
		preSum, res int
		preSCount   = map[int]int{0: 1}
	)
	for _, v := range nums {
		preSum += v
		// NOTE: diff := k - preSum // makesure the order
		diff := k - preSum
		if count, ok := preSCount[diff]; ok {
			res += count
		}
		preSCount[preSum]++
	}
	return res
}

// @leet end

