from typing import Optional
from type import TreeNode
from collections import deque


# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        q = deque([root])
        prev = root
        while q:
            top = q.popleft()
            if top is not None:
                if prev is None:
                    return False
                q.append(top.left)
                q.append(top.right)
            prev = top  # NOTE: prev is not inside queue. but maintain the order
            ## NOTE: order is happend before popleft
        return True


# @leet end

