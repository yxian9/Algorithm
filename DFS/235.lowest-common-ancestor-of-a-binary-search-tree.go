package leetcode

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil {
    return root
  } // will never happend
  if root.Val > p.Val && root.Val > q.Val {
    return lowestCommonAncestor(root.Left, p, q)
  }
  if root.Val < p.Val && root.Val < q.Val {
    return lowestCommonAncestor(root.Right, p, q)
  }
  return root // cant be nil
}
