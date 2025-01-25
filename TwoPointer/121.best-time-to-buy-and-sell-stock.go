package leetcode

// @leet start
func maxProfit(prices []int) int {
	curMin, best := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < curMin {
			curMin = prices[i]
		}
		best = max(best, prices[i]-curMin)
	}
	return best
}

// @leet end

