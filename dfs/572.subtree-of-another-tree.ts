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

function check(t1: TreeNode | null, t2: TreeNode | null): boolean {
  if (t1 === null && t2 === null) return true;
  if (t1 === null || t2 === null) return false;
  if ( t1.val === t2.val && check(t1.left, t2.left) && check(t1.right, t2.right)){
    return true
  }
  return false
}
function isSubtree(root: TreeNode | null, subRoot: TreeNode | null): boolean {
  if ( root === null) return false
  if (check(root, subRoot)) return true
  return isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot)
}
// @leet end

