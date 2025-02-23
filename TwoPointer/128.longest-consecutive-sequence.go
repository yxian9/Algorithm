package leetcode

// @leet start
func longestConsecutive(nums []int) int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	ret := 0
	for x := range freq {
		if freq[x-1] > 0 {
			continue
		}
		y := x + 1
		for freq[y] > 0 {
			y++
		}
		ret = max(ret, y-x)
	}
	return ret
}

// @leet end

