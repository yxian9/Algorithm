package leetcode

import "slices"

// @leet start
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var (
		res  [][]int
		l, r int
	)
	for idx, v := range nums {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		l, r = idx+1, len(nums)-1
		for l < r {
			nl, nr := nums[l], nums[r]
			sum := v + nl + nr
			if sum == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else { // sum < 0
				l++
			}
		}
	}
	return res
}

// @leet end

func threeSum_2(nums []int) [][]int {
	slices.Sort(nums)
	var (
		res  [][]int
		l, r int
	)
	for idx, v := range nums {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		l, r = idx+1, len(nums)-1
		// v + nums[l]+ nums[r] == 0
		for l < r {
			nl, nr := nums[l], nums[r]
			sum := v + nl + nr
			if sum == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				// for l+1 < r && nums[l+1] == nums[l] { // classic update idx and value
				// 	l++
				// }
				for l < r && nums[l] == nl {
					l++
				}
				for l < r && nums[r] == nr {
					r--
				}
			} else if sum > 0 {
				r--
			} else { // sum < 0
				l++
			}
		}
	}
	return res
}

