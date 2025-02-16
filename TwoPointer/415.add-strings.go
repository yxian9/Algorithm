package leetcode

import (
	"slices"
	"strings"
)

// @leet start
func addStrings(num1 string, num2 string) string {
	var s strings.Builder
	i1, i2 := len(num1)-1, len(num2)-1
	var n1, n2, carry int
	for i1 >= 0 || i2 >= 0 || carry != 0 {
		if i1 >= 0 {
			n1 = int(num1[i1]) - '0'
		} else {
			n1 = 0
		}
		if i2 >= 0 {
			n2 = int(num2[i2] - '0')
		} else {
			n2 = 0
		}
		sum := n1 + n2 + carry
		cur := sum % 10
		s.WriteRune(rune(cur) + '0')
		carry = sum / 10
		i1--
		i2--
	}
	// if carry > 0 {
	// 	s.WriteRune(rune(carry + '0'))
	// }
	ret := []rune(s.String())
	slices.Reverse(ret)
	return string(ret)
}

// @leet end

func addStrings(num1 string, num2 string) string {
	var (
		l1, l2 = len(num1) - 1, len(num2) - 1
		res    []byte
		carry  = 0
		total  = 0
	)
	for l1 >= 0 || l2 >= 0 {
		if l1 < 0 {
			total = int(num2[l2]-'0') + carry
		}
		if l2 < 0 {
			total = int(num1[l1]-'0') + carry
		}
		if l1 >= 0 && l2 >= 0 {
			total = int(num1[l1]-'0') + int(num2[l2]-'0') + carry
		}
		carry = total / 10
		res = append(res, byte(total%10+'0'))
		l1--
		l2--
	}
	if carry == 1 {
		res = append(res, byte(1+'0'))
	}
	slices.Reverse(res)
	return string(res)
}



