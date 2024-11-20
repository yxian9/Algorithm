package leetcode

import "math"

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfsValidBst(node *TreeNode, curMin, curMax int) bool {
  if node == nil {
    return true
  }
  x := node.Val
  if x >= curMax || x <= curMin {
    return false
  }
  return dfsValidBst(node.Left, curMin, x) && dfsValidBst(node.Right, x, curMax)
}

func isValidBST(root *TreeNode) bool {
  return dfsValidBst(root, math.MinInt, math.MaxInt)
}
// @leet end
