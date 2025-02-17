from typing import List


# @leet start
class Solution:
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        nr = len(mat)
        nc = len(mat[0])
        diagL = [[] for _ in range(nr + nc - 1)]

        for r in range(nr):
            for c in range(nc):
                diagL[r + c].append(mat[r][c])
        res = []
        for i, row in enumerate(diagL):
            if i % 2 == 0:
                row.reverse()
            res.extend(row)
        return res


# @leet end
