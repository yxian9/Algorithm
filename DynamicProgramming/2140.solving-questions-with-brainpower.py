from typing import List, Optional
from collections import defaultdict


# @leet start
class Solution:
    def mostPoints(self, questions: List[List[int]]) -> int:
        n = len(questions)
        cache = [0] * n
        for i in reversed(range(n)):
            p, cost = questions[i]
            cache[i] = max(
                cache[i + 1] if i + 1 < n else 0,
                p + (cache[i + 1 + cost] if i + 1 + cost < n else 0),
            )
        return cache[0]


# @leet end
def mostPoints2(self, questions: List[List[int]]) -> int:
    n = len(questions)
    cache = [0] * n

    def dfs(i):
        if i >= n:
            return 0
        if cache[i]:
            return cache[i]
        p, cost = questions[i]
        cache[i] = max(
            dfs(i + 1),
            # NOTE: not just advance to i+cost, but skip i+cost
            p + dfs(i + 1 + cost),
        )
        return cache[i]

    return dfs(0)

