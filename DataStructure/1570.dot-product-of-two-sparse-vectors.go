package leetcode

// @leet start
type SparseVector struct {
	nonZero map[int]int
}

func Constructor(nums []int) SparseVector {
	nonZero := map[int]int{}
	for i, v := range nums {
		if v != 0 {
			nonZero[i] = v
		}
	}
	return SparseVector{
		nonZero: nonZero,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	result := 0
	for i, v := range this.nonZero {
		result += v * vec.nonZero[i]
	}
	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
// @leet end

