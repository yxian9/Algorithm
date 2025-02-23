from collections import deque
from typing import List, Self
from math import inf

# @leet start
from dataclasses import dataclass


@dataclass(frozen=True)
class pt:
    r: int
    c: int

    def move(self, d: Self):
        return pt(self.r + d.r, self.c + d.c)


class Solution:
    dirs = (
        pt(0, 1),
        pt(0, -1),
        pt(1, 0),
        pt(-1, 0),
    )

    def minCost(self, grid: List[List[int]]) -> int:
        nr, nc = len(grid), len(grid[0])
        res = [[inf] * nc for _ in range(nr)]
        res[0][0] = 0

        def inside(p: pt):
            return 0 <= p.r < nr and 0 <= p.c < nc

        queue = deque([pt(0, 0)])

        while queue:
            p = queue.popleft()

            for i, d in enumerate(self.dirs, 1):
                np = p.move(d)
                cost = 0 if i == grid[p.r][p.c] else 1
                if inside(np) and res[p.r][p.c] + cost < res[np.r][np.c]:
                    res[np.r][np.c] = res[p.r][p.c] + cost
                    if cost == 0:
                        queue.appendleft(np)
                    else:
                        queue.append(np)

        return int(res[nr - 1][nc - 1])


# @leet end

