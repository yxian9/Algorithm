package leetcode

// @leet start
func checkSubarraySum(nums []int, k int) bool {
	var (
		preSum int
		m      = map[int]int{0: -1}
	)
	for i, v := range nums {
		preSum += v
		remain := preSum % k
		if preIdx, ok := m[remain]; ok {
			if i-preIdx >= 2 {
				return true
			}
		} else {
			m[remain] = i
		}
		// NOTE: m[remain] = i // need max length, avoid overwrite
	}
	return false
}

// @leet end

