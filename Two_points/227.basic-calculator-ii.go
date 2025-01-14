package leetcode

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

