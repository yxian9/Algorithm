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
