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
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        res = 0

        def dfs(n: TreeNode | None, s: str):
            if n is None:  # NOTE: required , consider node with only one child
                return
            nonlocal res
            if n.left is None and n.right is None:
                res += int(s + str(n.val))
                return
            s += str(n.val)
            dfs(n.left, s)
            dfs(n.right, s)

        dfs(root, "")
        return res


# @leet end

