package leetcode

import "strconv"

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var ret int
	var dfs func(node *TreeNode, cur string)
	dfs = func(node *TreeNode, cur string) {
		if node == nil {
			// always need, consider node only had left, the right will hit nil
			// need skip right nil node.
			return
		}
		curString := strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			num, _ := strconv.Atoi(cur + curString)
			ret += num
			return
		}
		dfs(node.Left, cur+curString)
		dfs(node.Right, cur+curString)
	}
	dfs(root, "")
	return ret
}

// @leet end

