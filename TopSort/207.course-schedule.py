from typing import List, Optional
from collections import defaultdict, deque


# @leet start
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        adj = defaultdict(list)
        indegree = {i: 0 for i in range(numCourses)}
        for dst, src in prerequisites:
            indegree[dst] += 1
            adj[src].append(dst)
        q = deque()
        for src, n in indegree.items():
            if n == 0:
                q.append(src)
        couse_take = 0
        while q:
            src = q.popleft()
            couse_take += 1
            for dst in adj[src]:
                indegree[dst] -= 1
                if indegree[dst] == 0:
                    q.append(dst)
        return couse_take == len(indegree)


# @leet end
