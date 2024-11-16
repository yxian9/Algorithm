package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	target := 0
	var dfs func(node *TreeNode, mx int)
	dfs = func(node *TreeNode, mx int) {
		if node == nil {
			return
		}
		if node.Val >= mx {
			target++
		}
		curMx := max(node.Val, mx)
		dfs(node.Left, curMx)
		dfs(node.Right, curMx)
	}
	dfs(root, math.MinInt)
	return target
}

// @leet end
