package leetcode

import "math"

// @leet start

// use a acount to check if all covered
// func allCovered(freq map[rune]int) bool {
// 	for _, v := range freq {
// 		if v > 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func minWindow(s string, t string) string {
	freq := map[rune]int{}
	for _, r := range t {
		freq[r]++
	}
	var bestL, bestR, l, r int
	minLen, count := math.MaxInt, len(t)
	for ; r < len(s); r++ {
		curR := rune(s[r])
		// move r until all covered
		if _, ok := freq[curR]; !ok {
			continue
		} else {
			freq[curR]--
			if freq[curR] >= 0 {
				count--
			}
			if count > 0 {
				continue
			}
			// if !allCovered(freq) {
			// 	continue
			// }
		}

		// allCovered, try to shirk l side
		for {
			lrune := rune(s[l])
			if _, ok := freq[lrune]; !ok {
				l++
			} else {
				if freq[lrune]+1 > 0 {
					break
				}
				freq[lrune]++
				l++
			}
		}

		curLen := r - l + 1
		if curLen < minLen { // need <=
			minLen = curLen
			bestL, bestR = l, r
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return s[bestL : bestR+1]
}

// @leet end
// original approach
func allCovered(freq map[rune]int) bool {
	for _, v := range freq {
		if v > 0 {
			return false
		}
	}
	return true
}

func minWindow2(s string, t string) string {
	freq := map[rune]int{}
	for _, r := range t {
		freq[r]++
	}
	var bestL, bestR, l, r int
	minLen := len(s)
	for ; r < len(s); r++ {
		curR := rune(s[r])
		// move r until all covered
		if _, ok := freq[curR]; !ok {
			continue
		} else {
			freq[curR]--
			if !allCovered(freq) {
				continue
			}
		}

		// allCovered, try to shirk l side
		for {
			lrune := rune(s[l])
			if _, ok := freq[lrune]; !ok {
				l++
			} else {
				if freq[lrune]+1 > 0 {
					break
				}
				freq[lrune]++
				l++
			}
		}

		curLen := r - l + 1
		if curLen <= minLen { // need <=
			minLen = curLen
			bestL, bestR = l, r
		}
	}
	if !allCovered(freq) {
		return ""
	}
	return s[bestL : bestR+1]
}

// use array
func minWindow(s string, t string) string {
	freq := [128]int{}    // can directly use array as char freq
	for _, r := range t { // up vote
		freq[r]++
	}
	var bestL, bestR, l, r int
	minLen, count := math.MaxInt, len(t)
	for ; r < len(s); r++ {
		curR := s[r]
		// move r until all covered
		freq[curR]--
		if freq[curR] >= 0 {
			count--
		}
		// if count > 0 {
		// 	continue
		// } // logic can move to next

		// if !allCovered(freq) {
		// 	continue
		// }

		// allCovered, try to shirk l side
		// for {
		for count == 0 {
			curLen := r - l + 1
			if curLen < minLen { // need <=
				minLen = curLen
				bestL, bestR = l, r
			}
			lrune := s[l]
			// if freq[lrune]+1 > 0 {
			//     count++
			// 	// break
			// }
			freq[lrune]++
			l++
			if freq[lrune] > 0 {
				count++
			}
		}

	}
	if minLen == math.MaxInt {
		return ""
	}
	return s[bestL : bestR+1]
}

// use a acount to check if all covered
// func allCovered(freq map[rune]int) bool {
// 	for _, v := range freq {
// 		if v > 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func minWindow(s string, t string) string {
	var (
		freq          [128]int
		l, bestl      int
		minLen, count = -1, len(t)
	)
	for _, b := range t {
		freq[b]++
	}
	for r, b := range s {
		freq[b]--
		if freq[b] >= 0 {
			count--
		}
		for count == 0 { // all freq [t] <= 0
			curLen := r - l + 1
			if minLen == -1 || curLen < minLen { // can be <= minLen does not matter
				minLen, bestl = curLen, l
			}
			lb := s[l] // pop left byte
			freq[lb]++
			l++
			if freq[lb] > 0 { // each byte is a single trigger point
				count++
			}
		}

	}
	if minLen == -1 {
		return ""
	}

	return s[bestl : bestl+minLen]
}
