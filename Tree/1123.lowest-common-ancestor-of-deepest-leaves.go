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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		l, lHeight := dfs(node.Left)
		r, rHeight := dfs(node.Right)
		if lHeight == rHeight {
			return node, lHeight + 1
		}
		if lHeight > rHeight {
			return l, lHeight + 1
		}
		return r, rHeight + 1
	}
	ret, _ := dfs(root)
	return ret
}

// @leet end

