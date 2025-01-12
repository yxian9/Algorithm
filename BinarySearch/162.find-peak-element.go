package leetcode

// @leet start
func findPeakElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[length-1] > nums[length-2] {
		return length - 1
	}
	l, r, mid := 1, length-2, -1
	for l <= r {
		mid = l + (r-l)>>1
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return mid
}

// @leet end

