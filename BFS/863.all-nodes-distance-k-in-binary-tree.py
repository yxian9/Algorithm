from collections import deque
from typing import List
from type import TreeNode

# @leet start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, k: int) -> List[int]:
        m = {}

        def dfs(n, p, np):
            if n is None:
                return
            np[n] = p
            dfs(n.left, n, np)
            dfs(n.right, n, np)

        dfs(root, None, m)
        q = deque([target])
        seen, step = {target}, 0
        while len(q) > 0:
            if step == k:
                break
            step += 1
            for _ in range(len(q)):
                cur = q.popleft()
                for v in [cur.left, cur.right, m[cur]]:
                    if v is not None and v not in seen:
                        q.append(v)
                        seen.add(v)

        return [v.val for v in q]


# @leet end

