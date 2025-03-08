from typing import List
from collections import defaultdict


# @leet start
class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        rowm = defaultdict(set)
        colm = defaultdict(set)
        subm = defaultdict(set)

        def valid(r, c, s):
            for i in range(9):
                if board[r][i] == s:
                    return False
                if board[i][c] == s:
                    return False
                if board[(r // 3) * 3 + i // 3][(c // 3) * 3 + i % 3] == s:
                    return False
            return True

        for i in range(9):
            for j in range(9):
                ch = board[i][j]
                if ch != ".":
                    rowm[i].add(ch)
                    colm[j].add(ch)
                    subm[i // 3, j // 3].add(ch)

        def dfs(r, c) -> bool:
            if c == 9:
                return dfs(r + 1, 0)
            if r == 9:
                return True
            s = board[r][c]
            if s != ".":
                return dfs(r, c + 1)

            for i in range(1, 10):
                s = str(i)
                if s in rowm[r] or s in colm[c] or s in subm[r // 3, c // 3]:
                    continue
                # if not valid(r, c, s):
                #     continue

                board[r][c] = s
                rowm[r].add(s)
                colm[c].add(s)
                subm[r // 3, c // 3].add(s)

                if dfs(r, c + 1):
                    return True

                board[r][c] = "."
                rowm[r].remove(s)
                colm[c].remove(s)
                subm[r // 3, c // 3].remove(s)

            return False

        dfs(0, 0)


# @leet end

