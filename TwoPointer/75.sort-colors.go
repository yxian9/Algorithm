package leetcode

// @leet start
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i < len(nums); {
		v := nums[i]
		if r > i && v == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
			continue
		}
		if l < i && v == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			continue
		}
		i++
	}
}

// @leet end

