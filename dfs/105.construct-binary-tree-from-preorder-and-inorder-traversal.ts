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

function buildTree(preorder: number[], inorder: number[]): TreeNode | null {
  const inorderM = new Map<number, number>();
  for (const [idx, val] of inorder.entries()) {
    inorderM.set(val, idx);
  }

  function dfsBuild(
    preStart: number,
    inStart: number,
    length: number,
  ): TreeNode | null {
    if (length === 0) return null;
    const rootVal = preorder[preStart];
    const inorderNeedle = inorderM.get(rootVal)!;
    const leftSize = inorderNeedle - inStart;
    const rightSize = length - leftSize - 1;

    return new TreeNode(
      rootVal,
      dfsBuild(preStart + 1, inStart, leftSize),
      dfsBuild(preStart + 1 + leftSize, inorderNeedle + 1, rightSize),
    );
  }
  return dfsBuild(0, 0, preorder.length);
}
// @leet end

