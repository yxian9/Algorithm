from typing import List


# @leet start
class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res, path = [], []

        def dfs(idx, total):
            if total >= target:
                if total == target:
                    res.append(path[:])
                return
            for i in range(idx, len(candidates)):
                if i > idx and candidates[i] == candidates[i -1]:
                    continue
                v = candidates[i]
                path.append(v)
                dfs(ix + 1, total + v)
                path.pop()

        dfs(0, 0)
        return res


# @leet end

