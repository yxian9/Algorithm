from typing import Optional
from type import TreeNode


# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def lcaDeepestLeaves(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def dfs(node):
            if node is None:
                return node, 0
            l, dep_l = dfs(node.left)
            r, dep_r = dfs(node.right)
            if dep_l == dep_r:
                return node, dep_r + 1
            if dep_r > dep_l:
                return r, dep_r + 1
            return l, dep_l + 1

        ret, _ = dfs(root)
        return ret


# @leet end
