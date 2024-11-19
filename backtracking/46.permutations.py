from typing import List
# @leet start


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        seen = [False] * n
        res, path = [], []

        def dfs(level):
            if level == n:
                res.append(path[:])
                return
            for i, v in enumerate(nums):
                if seen[i]: ## seen is using the i, not level
                    continue
                path.append(v)
                seen[i] = True
                dfs(level + 1)
                path.pop()
                seen[i] = False

        dfs(0)
        return res


# @leet end

