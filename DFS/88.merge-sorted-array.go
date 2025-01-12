package leetcode

// @leet start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
			continue
		} else {
			nums1[i] = nums2[p2]
			p2--
			continue
		}
	}
}

// @leet end
func mergetest(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if p1 < 0 {
			for j := range p2 + 1 {
				nums1[j] = nums2[j]
			}
			break
		}
		if nums2[p2] >= nums1[p1] {
			nums1[i] = nums2[p2]
			p2--
			continue
		} else {
			nums1[i] = nums1[p1]
			p1--
			continue
		}
	}
}

