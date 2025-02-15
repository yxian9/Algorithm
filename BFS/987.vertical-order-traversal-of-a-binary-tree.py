from typing import DefaultDict, List, Optional
from type import TreeNode
from collections import defaultdict, deque


# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from dataclasses import dataclass


# NOTE: use dataclass as struct to simplify code
@dataclass(slots=True)
class item:
    col: int
    row: int
    node: TreeNode


class Solution:
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        maxC = minC = 0
        colM: DefaultDict[int | str, list[item]] = defaultdict(list)
        q = deque([item(0, 0, root)])
        while len(q) > 0:
            cur = q.popleft()
            maxC = max(maxC, cur.col)
            minC = min(minC, cur.col)
            colM[cur.col].append(cur)
            if cur.node.left is not None:
                q.append(item(cur.col - 1, cur.row + 1, cur.node.left))
            if cur.node.right is not None:
                q.append(item(cur.col + 1, cur.row + 1, cur.node.right))

        res = []
        for i in range(minC, maxC + 1):
            cur = colM[i]
            cur.sort(key=lambda x: (x.row, x.node.val))
            res.append([v.node.val for v in cur])
        return res


# @leet end
