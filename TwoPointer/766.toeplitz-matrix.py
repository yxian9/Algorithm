from typing import List


# @leet start
class Solution:
    def isToeplitzMatrix(self, matrix: List[List[int]]) -> bool:
        nr, nc = len(matrix), len(matrix[0])

        def check(r, c):
            n = matrix[r][c]
            while r < nr and c < nc:
                if matrix[r][c] != n:
                    return False
                r += 1
                c += 1
            return True

        for r in range(nr):
            if not check(r, 0):
                return False
        for c in range(nc):
            if not check(0, c):
                return False
        return True


# @leet end

