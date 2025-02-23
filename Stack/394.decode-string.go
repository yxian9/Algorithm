package leetcode

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// @leet start
func decodeString(s string) string {
	ret := []rune{}
	for _, r := range s {
		if r == ']' { // NOTE: ] never get put in
			var ct, pattern []rune

			for len(ret) > 0 && ret[len(ret)-1] != '[' {
				pattern = append(pattern, ret[len(ret)-1])
				ret = ret[:len(ret)-1]
			}
			slices.Reverse(pattern)
			ret = ret[:len(ret)-1]

			for len(ret) > 0 && unicode.IsDigit(ret[len(ret)-1]) {
				ct = append(ct, ret[len(ret)-1])
				ret = ret[:len(ret)-1]
			}
			slices.Reverse(ct)
			count, _ := strconv.Atoi(string(ct))
			for range count {
				ret = append(ret, pattern...)
			}
		} else {
			ret = append(ret, r)
		}
	}

	return string(ret)
}

// @leet end
// NOTE: too smart
func decodeString2(s string) string {
	nums, codes := []int{}, []string{}
	num, code := "", ""
	for _, c := range s {
		if unicode.IsNumber(c) {
			num += string(c)
		} else if unicode.IsLetter(c) {
			code += string(c)
		} else if c == '[' {
			n, _ := strconv.Atoi(num)
			nums, codes = append(nums, n), append(codes, code)
			num, code = "", ""
			fmt.Printf("num %#v \n", nums)
			fmt.Printf("code %#v \n", codes)
		} else if c == ']' {
			fmt.Println("build")
			fmt.Printf("num %#v \n", nums)
			fmt.Printf("code %#v \n", codes)
			n, c := nums[len(nums)-1], codes[len(codes)-1]
			code = c + strings.Repeat(code, n)
			fmt.Println("res", code)
			nums, codes = nums[:len(nums)-1], codes[:len(codes)-1]
		}
	}
	return code
}

