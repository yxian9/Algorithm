package leetcode

import "strconv"

// @leet start
func maximumSwap(num int) int {
	input := []rune(strconv.Itoa(num))
	numIdx := make([]int, 10)
	// record the last freq of each integer
	for i, r := range input {
		numIdx[int(r-'0')] = i
	}

	for i, r := range input {
		// if an large integer show in later
		for v := 9; v > int(r-'0'); v-- {
			if laterIdx := numIdx[v]; laterIdx > i {
				input[laterIdx], input[i] = input[i], input[laterIdx]
				res, _ := strconv.Atoi(string(input))
				return res
			}
		}
	}
	return num
}

// @leet end

