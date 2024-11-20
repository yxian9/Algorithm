from typing import Optional
from type import TreeNode
from math import inf


# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def dfs(self, node, curMin: int | float, curMax: int | float) -> bool:
        if node is None:
            return True
        x = node.val
        if x >= curMax or x <= curMin:
            return False
        return self.dfs(node.left, curMin, x) and self.dfs(node.right, x, curMax)

    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.dfs(root, -inf, inf)


# @leet end

