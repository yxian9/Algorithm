import { TreeNode } from "./type.ts";

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

function lcaDeepestLeaves(root: TreeNode | null): TreeNode | null {
  function dfs(node: TreeNode | null): [number, TreeNode | null] {
    if (node === null) {
      return [0, null];
    }
    const [l, lca] = dfs(node.left);
    const [r, rca] = dfs(node.right);
    if (l === r) {
      return [l + 1, node];
    }
    if (l > r) {
      return [l + 1, lca];
    }
    return [r + 1, rca];
  }
  const [_, ret] = dfs(root);
  return ret;
}
// @leet end

