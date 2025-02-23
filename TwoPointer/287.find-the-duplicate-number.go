package leetcode

// @leet start
func findDuplicate(nums []int) int {
	var slow, fast int
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			fast = 0
			for {
				slow, fast = nums[slow], nums[fast]
				if slow == fast {
					return slow
				}
			}
		}
	}
}

// @leet end
