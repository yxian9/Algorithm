package leetcode

// @leet start
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	ret := [][]int{}

	cur := lower
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] > cur {
			ret = append(ret, []int{cur, nums[idx] - 1})
		}
		cur = nums[idx] + 1 // big step
	}
	// if nums[len(nums)-1] < upper {
	if cur <= upper {
		ret = append(ret, []int{cur, upper})
	}
	return ret
}

// @leet end

func findMissingRanges2(nums []int, lower int, upper int) [][]int {
	// loop over incorrect item
	ret := [][]int{}
	if len(nums) == 0 {
		return append(ret, []int{lower, upper})
	}
	// if nums[0] == lower && nums[len(nums)-1] == upper {
	// 	return ret
	// }

	i, idx := lower, 0
	var start int
	var findGap bool
	for i <= upper && idx < len(nums) {
		if i == nums[idx] {
			if findGap {
				findGap = false
				ret = append(ret, []int{start, i - 1})
			}
			idx++
			i++
			continue
		}
		if i < nums[idx] {
			if !findGap {
				start = i
			}
			findGap = true
			i++
		}
	}
	if nums[len(nums)-1] < upper {
		ret = append(ret, []int{nums[len(nums)-1] + 1, upper})
	}
	return ret
}

