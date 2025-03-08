from typing import List
from collections import defaultdict


# @leet start
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rowm = defaultdict(set)
        colm = defaultdict(set)
        subm = defaultdict(set)
        for r in range(9):
            for c in range(9):
                s = board[r][c]
                if s == ".":
                    continue
                if s in rowm or s in colm or s in subm[r // 3, c // 3]:
                    return False
                rowm[r].add(s)
                colm[c].add(s)
                subm[r // 3, c // 3].add(s)

        return True


# @leet end

