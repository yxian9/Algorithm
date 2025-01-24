package leetcode

// @leet start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if idx, ok := m[diff]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
    
	return nil
}

// @leet end
