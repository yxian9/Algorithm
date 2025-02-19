package leetcode

// @leet start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]bool{}
	for i, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
		if i >= k { // order matters!
			delete(m, nums[i-k])
		}
	}

	return false
}

// @leet end

func containsNearbyDuplicate4(nums []int, k int) bool {
	m := map[int]bool{}
	for i, v := range nums[k:] { // NOTE: not work. k may be large than
		if m[v] {
			return true
		}
		m[v] = true
		if i >= k { // NOTE: remove left item, so in the next round, it will not be check
			delete(m, nums[i-k])
		}
	}

	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	for r, v := range nums {
		for l := r - k; l < r; l++ { // NOTE: n^2, not O(n)
			if l >= 0 && nums[l] == v {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate3(nums []int, k int) bool {
	m := map[int]int{}
	for idx, v := range nums {
		if i, ok := m[v]; ok && idx-i <= k {
			return true
		}
		m[v] = idx
	}
	return false
}

