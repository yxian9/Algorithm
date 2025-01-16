package leetcode

// @leet start
func findKthPositive(arr []int, k int) int {
	l, r := 1, 0
	for r < len(arr) {
		if l == arr[r] {
			r++
		} else {
			k--
			if k == 0 {
				return l
			}
		}
		l++
	}
	// run out of r, last l is r lastItem +1
	return l + k - 1
}

// func findKthPositive(arr []int, k int) int {
//     j := 1
//     for i := 0; i < len(arr); j++ {
//         if arr[i] != j {
//             k--
//         } else {
//             i++
//         }
//         if k == 0 {
//             return j
//         }
//     }
//     return j + k - 1
// }
// @leet end

