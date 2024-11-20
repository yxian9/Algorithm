package leetcode

// @leet start
var (
	res      []string
	path     []rune
	keyboard = map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
)

func letterDfs(digits string) {
	if len(path) == len(digits) {
		res = append(res, string(path))
		return
	}
	curRune := rune(digits[len(path)])
	for _, r := range keyboard[curRune] {
		path = append(path, r)
		letterDfs(digits)
		path = path[:len(path)-1]
	}
}

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	letterDfs(digits)
	return res
}
