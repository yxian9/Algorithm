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

candidates.sort(reverse=True)
        ret = []
        path =[]
        seen = set()
        def dfs(idx,total):
            if total == target:
                ret.append(path[:])
                return
            if total > target:
                return
            for i in range(idx, len(candidates)):
                if i >0 and candidates[i] == candidates[i-1] and  i-1 not in seen:
                    continue
                if i in seen:
                    continue
                seen.add(i)
                item = candidates[i]
                path.append(item)
                dfs(i, total + item)
                path.pop()
                seen.remove(i)
        dfs(0,0)
        return ret
