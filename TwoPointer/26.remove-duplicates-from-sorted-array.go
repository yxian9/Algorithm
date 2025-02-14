package leetcode

// @leet start
func removeDuplicates(nums []int) int {
	l := 1
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v != nums[i-1] {
			nums[l] = v
			l++
		}
	}
	return l
}

// @leet end

func removeDuplicates(nums []int) int {
	l, r := 1, 1
	for r < len(nums) {
		if nums[r] == nums[r-1] {
			r++
			continue
		}
		nums[l] = nums[r]
		l++
		r++
	}
	return l
}

func removeDuplicates(nums []int) int {
	l, r := 1, 1
	for r < len(nums) {
		for r < len(nums) && nums[r] == nums[r-1] {
			r++
		} // need to keep check
		// when move slide window. just skip use one loop
		// avoid nested loop
		nums[l] = nums[r]
		l++
		r++
	}
	return l
}
