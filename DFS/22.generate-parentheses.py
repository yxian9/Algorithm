from typing import List


# @leet start
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def dfs(l, r, path):
            if l == r == n:
                res.append("".join(path))
                return
            if l < n:
                path.append("(")
                dfs(l + 1, r, path)
                path.pop()
            if l > r:
                path.append(")")
                dfs(l, r + 1, path)
                path.pop()

        dfs(0, 0, [])
        return res


# @leet end

