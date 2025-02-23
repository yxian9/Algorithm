from typing import List, Self


# @leet start
from dataclasses import dataclass


@dataclass(frozen=True)
class pt:
    r: int
    c: int

    def mv(self, d: Self):
        return pt(self.r + d.r, self.c + d.c)


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        res = [[0] * n for _ in range(n)]
        dirs = (pt(0, 1), pt(1, 0), pt(0, -1), pt(-1, 0))

        def inside(p: pt):
            return 0 <= p.r < n and 0 <= p.c < n

        step = 0
        cur = pt(0, 0)
        for k in range(1, n * n + 1):
            res[cur.r][cur.c] = k
            np = cur.mv(dirs[step % 4])
            if not inside(np) or res[np.r][np.c] != 0:
                step += 1
                np = cur.mv(dirs[step % 4])
            cur = np
        return res


# @leet end

