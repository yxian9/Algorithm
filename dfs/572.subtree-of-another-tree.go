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
func isSame(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val == t2.Val && isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right) {
		return true
	}
	return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if root == nil {
    return false
  }
  if isSame(root, subRoot){
    return true
  }
  if isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) {
    return true
  }
  // don't manuall check child node. use recursive
  // if isSame(root.Left, subRoot) {
  //   return true
  // }
  // if isSame(root.Right, subRoot) {
  //   return true
  // }
  return false
}

// @leet end

