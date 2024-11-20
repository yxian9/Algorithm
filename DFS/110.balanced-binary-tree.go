package leetcode

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getDepth(root.Left)
	if l == -1 {
		return -1
	}
	r := getDepth(root.Right)
	if r == -1 || abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}


func isBalanced(root *TreeNode) bool {
	return getDepth(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @leet end

