from typing import List
from collections import defaultdict


# @leet start
class DetectSquares:
    def __init__(self):
        self.pm = defaultdict(int)
        self.xym = defaultdict(set)

    def add(self, point: List[int]) -> None:
        x, y = point
        self.pm[x, y] += 1
        self.xym[x].add(y)

    def count(self, point: List[int]) -> int:
        ret = 0
        x, y = point
        for y2 in self.xym[x]:
            if y == y2:
                continue
            nx = x + y - y2, x + y2 - y
            for x2 in nx:
                ret += self.pm[x, y2] * self.pm[x2, y] * self.pm[x2, y2]
        return ret


# Your DetectSquares object will be instantiated and called as such:
# obj = DetectSquares()
# obj.add(point)
# param_2 = obj.count(point)
# @leet end

