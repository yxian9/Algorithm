from typing import List


# @leet start
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res, path, n = [], [], len(nums)

        def dfs(idx):
            res.append(path[:])
            for i in range(idx, n):
                path.append(nums[i])
                dfs(i + 1)
                path.pop()

        dfs(0)
        return res


# @leet end
class Solution2:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res, path, n = [], [], len(nums)

        def dfs(idx):
            if idx == n:
                res.append(path[:])
                return
            dfs(idx + 1)
            path.append(nums[idx])
            dfs(idx + 1)
            path.pop()

        dfs(0)
        return res


