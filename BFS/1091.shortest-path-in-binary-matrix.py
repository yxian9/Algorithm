from collections import deque

# @leet start
from dataclasses import dataclass
from typing import List, Self


@dataclass
class pt:
    r: int
    c: int

    def move(self, dir: Self):
        return pt(self.r + dir.r, self.c + dir.c)


class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        if grid[0][0] == 1:
            return -1
        nr, nc = len(grid), len(grid[0])
        queue: deque[pt] = deque([pt(0, 0)])
        grid[0][0] = 1
        dirs: list[pt] = []
        for i in range(-1, 2):
            for j in range(-1, 2):
                if i == 0 and j == 0:
                    continue
                dirs.append(pt(i, j))

        def get(p):
            return grid[p.r][p.c]

        def inside(p) -> bool:
            return p.r >= 0 and p.r < nr and p.c >= 0 and p.c < nc

        while len(queue) > 0:
            curP = queue.popleft()
            if curP.r == nr - 1 and curP.c == nc - 1:
                return get(curP)
            for dir in dirs:
                nextP = curP.move(dir)
                if inside(nextP) and get(nextP) == 0:
                    queue.append(nextP)
                    grid[nextP.r][nextP.c] = get(curP) + 1

        return -1


# @leet end

