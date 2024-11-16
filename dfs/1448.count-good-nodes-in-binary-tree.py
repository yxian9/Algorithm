from typing import Optional
from math import inf
from type import TreeNode

# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def dfs(self, root: TreeNode | None, mx: int | float) -> int:
        if root is None:
            return 0
        curMx = max(root.val, mx)
        l = self.dfs(root.left, curMx)
        r = self.dfs(root.right, curMx)
        if root.val >= mx:
            return l + r +  1
        return l + r
    def goodNodes(self, root: Optional[TreeNode]) -> int:
        return self.dfs(root, -inf)



        
# @leet end
