from typing import List, Self
from collections import deque


# @leet start
from dataclasses import dataclass


@dataclass(frozen=True)
class pt:
    r: int
    c: int

    def mv(self, d: Self):
        return pt(self.r + d.r, self.c + d.c)


class Solution:
    dirs = (
        pt(-1, 0),
        pt(1, 0),
        pt(0, -1),
        pt(0, 1),
    )

    def shortestBridge(self, grid: List[List[int]]) -> int:
        seen: set[pt] = set()
        nr, nc = len(grid), len(grid[0])
        # locate first land as seen

        def isInside(p: pt):
            return 0 <= p.r < nr and 0 <= p.c < nc

        def get(p: pt):
            return grid[p.r][p.c]

        def dfs(p: pt):
            if not isInside(p) or get(p) == 0 or p in seen:
                return
            seen.add(p)
            for d in self.dirs:
                dfs(p.mv(d))

        i, j = next((i, j) for i in range(nr) for j in range(nc) if grid[i][j])
        dfs(pt(i, j))
        # for r in range(nr):
        #     t_break = False
        #     for c in range(nc):
        #         if grid[r][c] == 1:
        #             dfs(pt(r, c))
        #             t_break = True
        #             break
        #     if t_break:
        #         break

        ## bfs
        q = deque(seen)
        ans = 0
        while q:
            for _ in range(len(q)):
                cur = q.popleft()
                # if cur not in seen and get(cur) == 1: # NOTE: again!
                #     return ans
                for d in self.dirs:
                    np = cur.mv(d)
                    if isInside(np) and np not in seen:
                        if get(np) == 1:
                            return ans
                        q.append(np)
                        seen.add(np)
            ans += 1

        return ans


# @leet end
