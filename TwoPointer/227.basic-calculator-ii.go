package leetcode

import "unicode"

func calculate_remove_previous(s string) int {
	var (
		ret, prev, num int // the prev is just a place holder for next ops. no need to add in the final return
		// do need append a + though
		prevOp = '+'
	)
	for _, r := range s + "+" { // use prev, still need additional ops to push last item
		if r == ' ' {
			continue
		}
		if unicode.IsNumber(r) {
			num = num*10 + int(r-'0')
			continue
		}
		// fond ops
		switch prevOp {
		case '+':
			ret, prev = ret+num, num
		case '-':
			ret, prev = ret-num, -num
		case '*':
			ret = ret - prev + prev*num
			prev = prev * num
		case '/':
			ret = ret - prev + prev/num
			prev = prev / num
		}
		// reset
		num, prevOp = 0, r
	}
	return ret
}

// with stack
func calculate3(s string) int {
	var num int
	prevOp := '+'
	s += "+"
	res := []int{}
	for _, r := range s {
		if r == ' ' {
			continue
		}
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
		} else {
			// found ops, process num, with prev
			// if preOps == +, can push to stack, add
			switch prevOp {
			case '+':
				res = append(res, num)
			case '-':
				res = append(res, -num)
			case '*':
				res[len(res)-1] = res[len(res)-1] * num
			case '/':
				res[len(res)-1] = res[len(res)-1] / num
			}

			num, prevOp = 0, r
		}
	}
	ret := 0
	for _, v := range res {
		ret += v
	}
	return ret
}

// lh, rh version
// lh, rh version
func calculate4(s string) int {
	var (
		ret, lh, rh int
		prevOp      = '+'
	)
	for _, r := range s + "+" { // + just finish last lh rh operator. still need add
		if r == ' ' {
			continue
		}
		if unicode.IsNumber(r) {
			rh = rh*10 + int(r-'0')
			continue
		}
		// fond ops
		switch prevOp {
		case '+':
			ret, lh = ret+lh, rh
		case '-':
			ret, lh = ret+lh, -rh
		case '*':
			lh *= rh
		case '/':
			lh /= rh
		}
		// reset
		rh, prevOp = 0, r
	}
	return ret + lh
}

func calculate_array_version(s string) int {
	var num int
	prevOp := '+'
	s += "+"
	res := []int{}
	for _, r := range s {
		if r == ' ' {
			continue
		}
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
		} else {
			// found ops, process num, with prev
			// if preOps == +, can push to stack, add
			switch prevOp {
			case '+':
				res = append(res, num)
			case '-':
				res = append(res, -num)
			case '*':
				res[len(res)-1] = res[len(res)-1] * num
			case '/':
				res[len(res)-1] = res[len(res)-1] / num
			}

			num, prevOp = 0, r
		}
	}
	ret := 0
	for _, v := range res {
		ret += v
	}
	return ret
}

// with out stack
func calculate2(s string) int {
	var res, prev, num int
	prevOp := '+'
	s += "+"
	for _, r := range s {
		if r == ' ' {
			continue
		}
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
		} else {
			// found ops, process num, with prev
			// if preOps == +, can push to stack, add
			switch prevOp {
			case '+':
				res += prev
				prev = num
			case '-':
				res += prev
				prev = -num
			case '*':
				prev *= num
			case '/':
				prev /= num
			}

			num, prevOp = 0, r
		}

	}
	return res + prev
}

// @leet start

func getNum(s string, idx int) (num, step int) {
	for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
		num *= 10
		num += int(s[idx] - '0')
		step++
		idx++
	}
	return num, step
}

// first verbose version
func calculate(s string) int {
	var nums []int
	idx := 0
	ops := '.'
	for idx < len(s) {
		r := rune(s[idx])
		if r == ' ' {
			idx++
			continue
		}
		if r >= '0' && r <= '9' {
			curNum, step := getNum(s, idx)
			idx += step
			// handle number
			if ops == '.' {
				nums = append(nums, curNum)
				continue
			} else {
				// do operation
				switch ops {
				case '+':
					nums = append(nums, curNum)
				case '-':
					nums = append(nums, -curNum)
				case '*':
					lastItem := nums[len(nums)-1]
					nums = nums[:len(nums)-1]
					nums = append(nums, lastItem*curNum)
				case '/':
					lastItem := nums[len(nums)-1]
					nums = nums[:len(nums)-1]
					nums = append(nums, lastItem/curNum)
				}
				// reset ops
				ops = '.'
			}
		}
		ops = r
		idx++
	}

	var res int
	for _, v := range nums {
		res += v
	}
	return res
}

// @leet end
