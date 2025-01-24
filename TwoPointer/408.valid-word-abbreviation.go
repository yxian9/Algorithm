package leetcode

import "strconv"

// @leet start
func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func validWordAbbreviation(word string, abbr string) bool {
	m, n := len(word), len(abbr)
	var wIdx, aIdx int
	for wIdx < m && aIdx < n {
		if word[wIdx] == abbr[aIdx] {
			wIdx++
			aIdx++
			continue
		}
		// not match
		// not number or start with 0
		if abbr[aIdx] <= '0' || abbr[aIdx] > '9' {
			return false
		}
		// build string
		ints := ""
		for aIdx < n && abbr[aIdx] >= '0' && abbr[aIdx] <= '9' {
			ints += string(abbr[aIdx])
			aIdx++
		}
		step, _ := strconv.Atoi(ints)
		wIdx += step
	}
	return wIdx == m && aIdx == n
}

// @leet end

func buildNum(s string, i int) (int, int) {
	res := 0
	for ; i < len(s); i++ {
		char := rune(s[i])
		if !isNum(char) {
			break
		}
		num := int(char - '0')
		res *= 10
		res += num
	}
	return res, i
}

func validWordAbbreviation2(word string, abbr string) bool {
	// expan abbr
	full := []rune{}
	idx := 0
	for idx < len(abbr) {
		r := rune(abbr[idx])
		if !isNum(r) {
			full = append(full, r)
			idx++
			continue
		}
		if r == '0' {
			return false
		}
		num, i := buildNum(abbr, idx)
		if len(full)+num > len(word) {
			return false
		}
		for range num {
			full = append(full, '*')
		}
		idx = i
	}
	if len(full) != len(word) {
		return false
	}
	for idx, v := range full {
		if v == '*' {
			continue
		}
		if v != rune(word[idx]) {
			return false
		}
	}
	return true
}
