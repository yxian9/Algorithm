from typing import List, Self

# @leet start
from dataclasses import dataclass


@dataclass(slots=True)
class pt:
    r: int
    c: int

    def move(self, d: Self):
        return pt(self.r + d.r, self.c + d.c)


class Solution:
    dirs = [
        pt(0, 1),
        pt(0, -1),
        pt(-1, 0),
        pt(1, 0),
    ]

    def largestIsland(self, grid: List[List[int]]) -> int:
        group_area: dict[int, int] = {0: 0}
        res = 0
        nr, nc = len(grid), len(grid[0])
        group = 2

        def get(p: pt):
            return grid[p.r][p.c]

        def inside(p: pt):
            return p.r >= 0 and p.r < nr and p.c >= 0 and p.c < nc

        def dfs(p: pt, g: int) -> int:
            if not inside(p) or get(p) != 1:
                return 0
            grid[p.r][p.c] = g
            area = 1
            for d in self.dirs:
                area += dfs(p.move(d), g)
            return area

        def findMax(p: pt):
            seen = set()
            cur = 1
            for d in self.dirs:
                ap = p.move(d)
                if inside(ap) and get(ap) not in seen:
                    cur += group_area[get(ap)]
                    seen.add(get(ap))
            return cur

        for r in range(nr):
            for c in range(nc):
                if grid[r][c] == 1:
                    area = dfs(pt(r, c), group)
                    res = max(area, res)
                    group_area[group] = area
                    group += 1

        for r in range(nr):
            for c in range(nc):
                if grid[r][c] == 0:
                    res = max(findMax(pt(r, c)), res)
        return res


# @leet end

