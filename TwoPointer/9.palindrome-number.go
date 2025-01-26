package leetcode

// @leet start
func isPalindrome(x int) bool {
	var (
		num = x
		res = 0
	)
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return x == res
}

// @leet end

