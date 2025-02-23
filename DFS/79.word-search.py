from typing import List, Self


# @leet start
from dataclasses import dataclass


@dataclass(frozen=True)
class pt:
    r: int
    c: int

    def move(self, d: Self):
        return pt(r=self.r + d.r, c=self.c + d.c)


class Solution:
    dirs = [
        pt(0, 1),
        pt(0, -1),
        pt(-1, 0),
        pt(1, 0),
    ]

    def exist(self, board: List[List[str]], word: str) -> bool:
        nr, nc = len(board), len(board[0])

        def get(p: pt):
            return board[p.r][p.c]

        def outSide(p: pt):
            return not (p.r >= 0 and p.r < nr and p.c >= 0 and p.c < nc)

        seen = set()

        def dfs(p: pt, step: int) -> bool:
            if outSide(p) or pt in seen or get(p) != word[step]:
                return False
            if step == len(word) - 1:
                return True
            seen.add(pt)
            for d in self.dirs:
                if dfs(p.move(d), step + 1):
                    return True
            seen.remove(pt)
            return False

        for r in range(nr):
            for c in range(nc):
                p = pt(r, c)
                if dfs(p, 0):
                    return True
        return False


# @leet end

