from typing import List
from collections import defaultdict
from math import inf


# @leet start
class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        def slope(p1, p2):
            x1, y1 = p1
            x2, y2 = p2
            if y1 == y2:
                return inf
            return (x1 - x2) / (y1 - y2)

        res = 0
        for i, p1 in enumerate(points):
            freq = defaultdict(int)
            for j, p2 in enumerate(points):
                if i == j:
                    continue
                temp = slope(p1, p2)
                freq[temp] += 1
                res = max(res, freq[temp])
        return res + 1


# @leet end

