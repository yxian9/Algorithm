from typing import List
from typing import NamedTuple


# @leet start
class Pt(NamedTuple):
    r: int
    c: int

    def mv(self, d: "Pt") -> "Pt":
        return Pt(self.r + d.r, self.c + d.c)


class Solution:
    dirs = [
        Pt(1, 0),
        Pt(-1, 0),
        Pt(0, -1),
        Pt(0, 1),
    ]

    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        ret = 0
        nr, nc = len(grid), len(grid[0])

        def inside(p: Pt):
            return 0 <= p.r < nr and 0 <= p.c < nc

        def dfs(p: Pt):
            if not inside(p) or grid[p.r][p.c] != 1:
                return 0
            ret = 1
            grid[p.r][p.c] = 2
            for d in self.dirs:
                ret += dfs(p.mv(d))
            return ret

        for r in range(nr):
            for c in range(nc):
                if grid[r][c] == 1:
                    p = Pt(r, c)
                    ret = max(ret, dfs(p))

        return ret


# @leet end

