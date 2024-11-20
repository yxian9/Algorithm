from typing import List


# @leet start
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res, path = [], []

        def dfs(i, total):
            if total >= target:
                if total == target:
                    res.append(path[:])
                return
            if i == len(candidates):
                return
            dfs(i + 1, total)
            path.append(candidates[i])
            dfs(i, total + candidates[i])
            path.pop()

        dfs(0, 0)
        return res


# @leet end
class Solution2:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res, path = [], []

        def dfs(idx, total):
            if total >= target:
                if total == target:
                    res.append(path[:])
                return
            for i in range(idx, len(candidates)):
                path.append(candidates[i])
                dfs(i, total + candidates[i])
                path.pop()

        dfs(0, 0)
        return res

