from typing import Optional

# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def getDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        l = self.getDepth(root.left)
        if l == -1:
            return -1
        r = self.getDepth(root.right)
        if abs(l - r) > 1 or r == -1:
            return -1
        return max(l, r) + 1
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        return self.getDepth(root) != -1

# @leet end
