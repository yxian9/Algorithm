package leetcode

// @leet start
func find(nums []int, target int, left bool) int {
	l, r, idx := 0, len(nums)-1, -1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			idx = mid
			if left {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{find(nums, target, true), find(nums, target, false)}
}

// @leet end

