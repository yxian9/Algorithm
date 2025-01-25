package leetcode

import (
	"slices"
	"strings"
)

// @leet start
func getNum(s string) (num int) {
	for _, r := range s {
		num = num*10 + int(r) - '0'
	}
	return num
}

func addOperators(num string, target int) []string {
	n := len(num)
	res := []string{}

	var dfsAdd func(idx, prev, cur int, buildS string)

	dfsAdd = func(idx, prev, cur int, buildS string) {
		if idx == n {
			if cur == target {
				res = append(res, buildS)
			}
		}
		for i := idx; i < n; i++ {
			nStr := num[idx : i+1]
			nNum := getNum(nStr)
			if num[idx] == '0' && i != idx { // idx not i
				break
			}
			if idx == 0 { // first index // no i
				dfsAdd(i+1, nNum, nNum, nStr)
			} else {
				dfsAdd(i+1, nNum, cur+nNum, buildS+"+"+nStr)
				dfsAdd(i+1, -nNum, cur-nNum, buildS+"-"+nStr)
				dfsAdd(i+1, prev*nNum, cur-prev+prev*nNum, buildS+"*"+nStr)
			}
		}
	}
	dfsAdd(0, 0, 0, "")
	return res
}

// @leet end

func process(op rune, lh int, b byte, pre int) (int, int) {
	rh := int(b) - '0'
	switch op {
	case '+':
		return lh + rh, rh
	case '-':
		return lh - rh, -rh
	case '*':
		return lh - pre + pre*rh, pre * rh
	default:
		return 0, 0
	}
}

func addOperators_previous_not_work(num string, target int) []string {
	var (
		n    = len(num)
		res  [][]rune
		ret  []string
		path []rune
		ops  = [3]rune{'+', '-', '*'}
	)

	var dfs func(idx, cur, prev int)
	dfs = func(idx, cur, prev int) {
		if idx == n {
			if cur == target {
				res = append(res, slices.Clone(path))
			}
			return
		}
		for _, r := range ops {
			cur, prev = process(r, cur, num[idx], prev)
			dfs(idx+1, cur, prev)
		}
	}
	dfs(1, int(num[0])-'0', int(num[0])-'0')
	for _, v := range res {
		var s strings.Builder
		s.WriteByte(num[0])
		for i, r := range v {
			s.WriteRune(r)
			s.WriteByte(num[i+1])
		}
		ret = append(ret, s.String())

	}
	return ret
}
