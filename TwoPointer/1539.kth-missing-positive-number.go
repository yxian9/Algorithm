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

package leetcode

// @leet start
func findKthPositive(arr []int, k int) int {
	l, r, idx := 0, len(arr)-1, -1
	for l <= r {
    mid := (l+r) >> 1
    if arr[mid] - (mid +1) < k {
      idx = mid // NOTE: binary serarch
      l = mid +1
    } else {
      r = mid -1
    }
	}
	return idx + 1 + k
}


func findKthPositive2(arr []int, k int) int {
	var missed int
	if arr[0] != 1 {
		missed = arr[0] - 1
	}
	if k <= missed {
		return k
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]+1 {
			miss := arr[i] - arr[i-1] - 1
			if missed+miss >= k {
				return arr[i-1] + k - missed
			} else {
				missed += miss
			}
		}
	}
	return arr[len(arr)-1] + k - missed
}


