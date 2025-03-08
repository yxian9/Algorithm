from typing import List, Optional
from collections import defaultdict
from type import TreeNode


# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def recoverTree(self, root: Optional[TreeNode]) -> None:
        if root is None:
            return
        prev = n1 = n2 = None
        # st = [root] # NOTE: not queue
        st = []
        cur = root
        while cur or st:
            while cur:
                st.append(cur)
                cur = cur.left
            cur = st.pop()
            if prev and prev.val > cur.val:
                if n1 is None:
                    n1 = prev
                    n2 = cur
                else:
                    n2 = cur
            prev = cur
            cur = cur.right

        n1.val, n2.val = n2.val, n1.val  # type: ignore


# @leet end
def recoverTree(root: Optional[TreeNode]) -> None:
    """
    Do not return anything, modify root in-place instead.
    """
    prev = n1 = n2 = None

    def dfs(node):
        if node is None:
            return
        nonlocal prev, n1, n2
        dfs(node.left)
        if prev and prev.val > node.val:
            if n1 is None:
                n1 = prev
                n2 = node
            else:
                n2 = node
                return
        prev = node
        dfs(node.right)

    dfs(root)
    n1.val, n2.val = n2.val, n1.val

