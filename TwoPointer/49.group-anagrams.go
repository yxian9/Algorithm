package leetcode

// @leet start
func groupAnagrams(strs []string) [][]string {
	var (
		freq  = map[[128]rune][]string{}
		parse = func(s string) (res [128]rune) {
			for _, r := range s {
				res[r]++
			}
			return res
		}
	)
	for _, s := range strs {
		key := parse(s)
		freq[key] = append(freq[key], s)
	}
	ret := make([][]string, 0, len(freq))
	for _, v := range freq {
		ret = append(ret, v)
	}

	return ret
}

// @leet end

