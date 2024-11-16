import { TreeNode } from "./type";
// @leet start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */
function getDepth(root: TreeNode | null): number {
  if (root == null) return 0;
  let l = getDepth(root.left);
  if (l == -1) return -1;
  let r = getDepth(root.right);
  if (r == -1 || Math.abs(l - r) > 1) return -1;
  return Math.max(l, r) + 1;
}
function isBalanced(root: TreeNode | null): boolean {
  return getDepth(root) !== -1
}
// @leet end

